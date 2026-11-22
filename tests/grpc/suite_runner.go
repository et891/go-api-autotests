package grpc

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/et891/go-api-autotests/metadata"
	"github.com/et891/go-api-autotests/tests"
)

var suiteRunner = tests.Runner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag(metadata.TagGRPC),
			axiom.WithMetaSuite(metadata.SuiteGRPC),
		),
	),
)

var usersRunner = suiteRunner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag(metadata.TagUsers),
			axiom.WithMetaFeature(metadata.FeatureUsers),
		),
	),
)

var authenticationRunner = suiteRunner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag(metadata.TagAuthentication),
			axiom.WithMetaFeature(metadata.FeatureAuthentication),
		),
	),
)
