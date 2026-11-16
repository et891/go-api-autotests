package tests

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
)

type axiomCourse struct {
	Title     string
	Published bool
}

var coursesRunner = testRunner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag("courses"),
		),
	),
)

func (s *AxiomSuite) TestCourseCanBeCreated() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("course can be created"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag("create"),
		))

	s.RunCase(testCase, func(cfg *axiom.Config) {
		expectedTitle := "Go API Autotests"

		course := axiomCourse{
			Title:     expectedTitle,
			Published: false,
		}

		assert.Equal(cfg.T(), expectedTitle, course.Title)
		assert.False(cfg.T(), course.Published)
		assert.Equal(
			cfg.T(),
			[]string{"axiom", "courses", "create"},
			cfg.Meta.Tags,
		)
	})
}

func (s *AxiomSuite) TestCourseCanBePublished() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("course can be published"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag("publish"),
		))

	s.RunCase(testCase, func(cfg *axiom.Config) {
		expectedTitle := "Go API Autotests"

		course := axiomCourse{
			Title:     expectedTitle,
			Published: false,
		}

		course.Published = true

		assert.True(cfg.T(), course.Published)
		assert.Equal(
			cfg.T(),
			[]string{"axiom", "courses", "publish"},
			cfg.Meta.Tags,
		)
	})
}
