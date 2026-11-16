package tests

import (
	"testing"

	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
)

type axiomCourse struct {
	Title     string
	Published bool
}

func TestAxiomCreateCourse(t *testing.T) {
	testRunner := axiom.NewRunner()
	testCase := axiom.NewCase(
		axiom.WithCaseName("course can be created"),
	)

	testRunner.RunCase(t, testCase, func(cfg *axiom.Config) {
		expectedTitle := "Go API Autotests"

		course := axiomCourse{
			Title:     expectedTitle,
			Published: false,
		}

		assert.Equal(cfg.T(), expectedTitle, course.Title)
		assert.False(cfg.T(), course.Published)
	})
}

func TestAxiomPublishCourse(t *testing.T) {
	testRunner := axiom.NewRunner()
	testCase := axiom.NewCase(
		axiom.WithCaseName("course can be published"),
	)

	testRunner.RunCase(t, testCase, func(cfg *axiom.Config) {
		expectedTitle := "Go API Autotests"

		course := axiomCourse{
			Title:     expectedTitle,
			Published: false,
		}

		course.Published = true

		assert.True(cfg.T(), course.Published)
	})
}
