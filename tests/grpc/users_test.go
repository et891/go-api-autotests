package grpc

import (
	"context"

	v1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/Nikita-Filonov/axiom"
	"github.com/et891/go-api-autotests/builders"
	"github.com/et891/go-api-autotests/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func (s *suite) TestCreateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("create user"),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		builder := builders.New(fake.New())
		user := builder.UserCreate()
		request := user.GRPCRequest()

		connection, err := grpc.NewClient(
			"localhost:9000",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		require.NoError(cfg.T(), err)
		defer connection.Close()

		client := v1.NewUsersServiceClient(connection)
		response, err := client.CreateUser(context.Background(), request)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetUser())
		require.NotEmpty(cfg.T(), response.GetUser().GetId())

		assert.Equal(cfg.T(), request.GetEmail(), response.GetUser().GetEmail())
		assert.Equal(cfg.T(), request.GetLastName(), response.GetUser().GetLastName())
		assert.Equal(cfg.T(), request.GetFirstName(), response.GetUser().GetFirstName())
		assert.Equal(cfg.T(), request.GetMiddleName(), response.GetUser().GetMiddleName())

		cfg.T().Logf("created user with ID %s", response.GetUser().GetId())
	})
}

func (s *suite) TestUpdateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("update user"),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		builder := builders.New(fake.New())
		user := builder.UserCreate()
		userRequest := user.GRPCRequest()

		connection, err := grpc.NewClient(
			"localhost:9000",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		require.NoError(cfg.T(), err)
		defer connection.Close()

		usersClient := v1.NewUsersServiceClient(connection)

		createdUser, err := usersClient.CreateUser(
			context.Background(),
			userRequest,
		)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), createdUser)
		require.NotNil(cfg.T(), createdUser.GetUser())
		require.NotEmpty(cfg.T(), createdUser.GetUser().GetId())

		authenticationClient :=
			v1.NewAuthenticationServiceClient(connection)

		authRequest := &v1.LoginRequest{
			Email:    user.Email,
			Password: user.Password,
		}

		loginResponse, err := authenticationClient.Login(
			context.Background(),
			authRequest,
		)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), loginResponse)
		require.NotNil(cfg.T(), loginResponse.GetToken())
		require.NotEmpty(
			cfg.T(),
			loginResponse.GetToken().GetAccessToken(),
		)

		update := builder.UserUpdate()
		updateRequest := update.GRPCRequest(createdUser.GetUser().GetId())

		authorizationMetadata := metadata.Pairs(
			"authorization",
			"Bearer "+loginResponse.GetToken().GetAccessToken(),
		)

		contextWithToken := metadata.NewOutgoingContext(
			context.Background(),
			authorizationMetadata,
		)

		response, err := usersClient.UpdateUser(
			contextWithToken,
			updateRequest,
		)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetUser())

		assert.Equal(cfg.T(), createdUser.GetUser().GetId(), response.GetUser().GetId())
		assert.Equal(cfg.T(), updateRequest.GetEmail(), response.GetUser().GetEmail())
		assert.Equal(cfg.T(), updateRequest.GetLastName(), response.GetUser().GetLastName())
		assert.Equal(cfg.T(), updateRequest.GetFirstName(), response.GetUser().GetFirstName())
		assert.Equal(cfg.T(), updateRequest.GetMiddleName(), response.GetUser().GetMiddleName())

		cfg.T().Logf("updated user with ID %s", response.GetUser().GetId())
	})
}

func (s *suite) TestGetUserMe() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("get my user"),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		builder := builders.New(fake.New())
		user := builder.UserCreate()
		userRequest := user.GRPCRequest()

		connection, err := grpc.NewClient(
			"localhost:9000",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		require.NoError(cfg.T(), err)
		defer connection.Close()

		usersClient := v1.NewUsersServiceClient(connection)
		authenticationClient := v1.NewAuthenticationServiceClient(connection)

		createdUser, err := usersClient.CreateUser(context.Background(), userRequest)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), createdUser)
		require.NotNil(cfg.T(), createdUser.GetUser())
		require.NotEmpty(cfg.T(), createdUser.GetUser().GetId())

		authRequest := &v1.LoginRequest{
			Email:    userRequest.GetEmail(),
			Password: userRequest.GetPassword(),
		}

		loginResponse, err := authenticationClient.Login(context.Background(), authRequest)

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), loginResponse)
		require.NotNil(cfg.T(), loginResponse.GetToken())
		require.NotEmpty(cfg.T(), loginResponse.GetToken().GetAccessToken())

		authorizationMetadata := metadata.Pairs(
			"authorization",
			"Bearer "+loginResponse.GetToken().GetAccessToken(),
		)

		contextWithToken := metadata.NewOutgoingContext(
			context.Background(),
			authorizationMetadata,
		)

		response, err := usersClient.GetMe(contextWithToken, &v1.Empty{})

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.NotNil(cfg.T(), response.GetUser())

		assert.Equal(cfg.T(), createdUser.GetUser().GetId(), response.GetUser().GetId())
		assert.Equal(cfg.T(), createdUser.GetUser().GetEmail(), response.GetUser().GetEmail())
		assert.Equal(cfg.T(), createdUser.GetUser().GetLastName(), response.GetUser().GetLastName())
		assert.Equal(cfg.T(), createdUser.GetUser().GetFirstName(), response.GetUser().GetFirstName())
		assert.Equal(cfg.T(), createdUser.GetUser().GetMiddleName(), response.GetUser().GetMiddleName())

		cfg.T().Logf("received current user with ID %s", response.GetUser().GetId())
	})
}
