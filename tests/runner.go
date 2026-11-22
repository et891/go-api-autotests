package tests

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/et891/go-api-autotests/metadata"
)

var Runner = axiom.NewRunner(
	axiom.WithRunnerMeta(
		axiom.WithMetaTags(metadata.TagAPI, metadata.TagRegression),
		axiom.WithMetaEpic(metadata.EpicAPI),
		axiom.WithMetaLayer(metadata.LayerE2E),
	),
)
