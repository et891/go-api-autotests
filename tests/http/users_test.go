package http

import (
	"net/http"

	"github.com/Nikita-Filonov/axiom"
	"github.com/et891/go-api-autotests/builders"
	"github.com/et891/go-api-autotests/clients/httpclients"
	"github.com/et891/go-api-autotests/config"
	"github.com/et891/go-api-autotests/fake"
	"github.com/et891/go-api-autotests/metadata"
	"github.com/et891/go-api-autotests/models"
	"github.com/et891/go-api-autotests/transport/httptransport"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (s *suite) TestCreateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("create user"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryCreateEntity),
			axiom.WithMetaSeverity(axiom.SeverityBlocker),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		settings, err := config.Load()
		require.NoError(cfg.T(), err)

		publicTransport := httptransport.NewPublic(cfg, settings.HTTP)
		publicUsersClient := httpclients.NewUsersClient(publicTransport)

		builder := builders.New(fake.New())
		user := builder.UserCreate()
		userRequest := user.HTTPRequest()

		createdUserResponse, err := publicUsersClient.Create(cfg.Context.Raw, userRequest)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, createdUserResponse.StatusCode)
		require.NotEmpty(cfg.T(), createdUserResponse.Data.User.ID)

		assert.Equal(cfg.T(), userRequest.Email, createdUserResponse.Data.User.Email)
		assert.Equal(cfg.T(), userRequest.LastName, createdUserResponse.Data.User.LastName)
		assert.Equal(cfg.T(), userRequest.FirstName, createdUserResponse.Data.User.FirstName)
		assert.Equal(cfg.T(), userRequest.MiddleName, createdUserResponse.Data.User.MiddleName)

		cfg.T().Logf("created user with ID %s", createdUserResponse.Data.User.ID)
	})
}

func (s *suite) TestUpdateUser() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("update user"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory(metadata.StoryUpdateEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		settings, err := config.Load()
		require.NoError(cfg.T(), err)

		publicTransport := httptransport.NewPublic(cfg, settings.HTTP)
		publicUsersClient := httpclients.NewUsersClient(publicTransport)
		authenticationClient := httpclients.NewAuthenticationClient(publicTransport)

		builder := builders.New(fake.New())
		user := builder.UserCreate()
		userRequest := user.HTTPRequest()

		createdUserResponse, err := publicUsersClient.Create(cfg.Context.Raw, userRequest)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, createdUserResponse.StatusCode)
		require.NotEmpty(cfg.T(), createdUserResponse.Data.User.ID)

		authRequest := models.LoginRequest{
			Email:    user.Email,
			Password: user.Password,
		}

		loginResponse, err := authenticationClient.Login(cfg.Context.Raw, authRequest)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, loginResponse.StatusCode)
		require.NotEmpty(cfg.T(), loginResponse.Data.Token.AccessToken)

		privateTransport := httptransport.NewPrivate(
			cfg,
			settings.HTTP,
			loginResponse.Data.Token.AccessToken,
		)
		privateUsersClient := httpclients.NewUsersClient(privateTransport)

		update := builder.UserUpdate()
		request := update.HTTPRequest()

		response, err := privateUsersClient.Update(
			cfg.Context.Raw,
			createdUserResponse.Data.User.ID,
			request,
		)

		require.NoError(cfg.T(), err)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode)

		assert.Equal(cfg.T(), createdUserResponse.Data.User.ID, response.Data.User.ID)
		assert.Equal(cfg.T(), *request.Email, response.Data.User.Email)
		assert.Equal(cfg.T(), *request.LastName, response.Data.User.LastName)
		assert.Equal(cfg.T(), *request.FirstName, response.Data.User.FirstName)
		assert.Equal(cfg.T(), *request.MiddleName, response.Data.User.MiddleName)

		cfg.T().Logf("updated user with ID %s", response.Data.User.ID)
	})
}

func (s *suite) TestGetUserMe() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("get user me"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag(metadata.TagSmoke),
			axiom.WithMetaStory(metadata.StoryGetEntity),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
		),
	)

	s.RunCase(testCase, func(cfg *axiom.Config) {
		builder := builders.New(fake.New())
		user := builder.UserCreate()
		createUserRequest := user.HTTPRequest()
		client := resty.New().
			SetBaseURL("http://localhost:8000/api/v1")

		var createdUser models.UserResponse

		createUserResponse, err := client.
			R().
			SetBody(createUserRequest).
			SetResult(&createdUser).
			Post("/users")

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), createUserResponse)
		require.Equal(cfg.T(), http.StatusOK, createUserResponse.StatusCode())
		require.NotEmpty(cfg.T(), createdUser.User.ID)

		loginRequest := models.LoginRequest{
			Email:    user.Email,
			Password: user.Password,
		}

		var loginResult models.LoginResponse

		loginResponse, err := client.
			R().
			SetBody(loginRequest).
			SetResult(&loginResult).
			Post("/authentication/login")

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), loginResponse)
		require.Equal(cfg.T(), http.StatusOK, loginResponse.StatusCode())
		require.NotEmpty(cfg.T(), loginResult.Token.AccessToken)

		var result models.UserResponse

		response, err := client.
			R().
			SetHeader("Authorization", "Bearer "+loginResult.Token.AccessToken).
			SetResult(&result).
			Get("/users/me")

		require.NoError(cfg.T(), err)
		require.NotNil(cfg.T(), response)
		require.Equal(cfg.T(), http.StatusOK, response.StatusCode())

		assert.Equal(cfg.T(), createdUser.User.ID, result.User.ID)
		assert.Equal(cfg.T(), createdUser.User.Email, result.User.Email)
		assert.Equal(cfg.T(), createdUser.User.LastName, result.User.LastName)
		assert.Equal(cfg.T(), createdUser.User.FirstName, result.User.FirstName)
		assert.Equal(cfg.T(), createdUser.User.MiddleName, result.User.MiddleName)

		cfg.T().Logf("received current user with ID %s", result.User.ID)
	})
}
