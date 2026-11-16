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
			axiom.WithMetaFeature("Courses"),
			axiom.WithMetaSuite("Course management"),
			axiom.WithMetaSubSuite("Lifecycle"),
			axiom.WithMetaLabel("owner", "courses-team"),
			axiom.WithMetaLabel("component", "courses-service"),
		),
	),
)

func (s *AxiomSuite) TestCourseCanBeCreated() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("course can be created"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory("Create course"),
			axiom.WithMetaSeverity(axiom.SeverityCritical),
			axiom.WithMetaTags("create", "smoke"),
			axiom.WithMetaIssue("AXIOM-201"),
			axiom.WithMetaTestCase("COURSES-001"),
		))

	s.RunCase(testCase, func(cfg *axiom.Config) {
		assert.Equal(cfg.T(), "Learning platform", cfg.Meta.Epic)
		assert.Equal(cfg.T(), "Courses", cfg.Meta.Feature)
		assert.Equal(cfg.T(), "Create course", cfg.Meta.Story)
		assert.Equal(cfg.T(), axiom.SeverityCritical, cfg.Meta.Severity)
		expectedTitle := "Go API Autotests"

		course := axiomCourse{
			Title:     expectedTitle,
			Published: false,
		}

		assert.Equal(cfg.T(), expectedTitle, course.Title)
		assert.False(cfg.T(), course.Published)
		assert.Equal(
			cfg.T(),
			[]string{"axiom", "courses", "create", "smoke"},
			cfg.Meta.Tags,
		)
		assert.Equal(cfg.T(), []string{"AXIOM-201"}, cfg.Meta.Issues)
		assert.Equal(cfg.T(), []string{"COURSES-001"}, cfg.Meta.TestCases)

		expectedLabels := map[string]string{
			"owner":     "courses-team",
			"component": "courses-service",
		}
		assert.Equal(cfg.T(), expectedLabels, cfg.Meta.Labels)
	})
}

func (s *AxiomSuite) TestCourseCanBePublished() {
	testCase := axiom.NewCase(
		axiom.WithCaseName("course can be published"),
		axiom.WithCaseMeta(
			axiom.WithMetaTags("publish", "regression"),
			axiom.WithMetaStory("Publish course"),
			axiom.WithMetaIssue("AXIOM-202"),
			axiom.WithMetaTestCase("COURSES-002"),
		))

	s.RunCase(testCase, func(cfg *axiom.Config) {
		assert.Equal(cfg.T(), "Learning platform", cfg.Meta.Epic)
		assert.Equal(cfg.T(), "Courses", cfg.Meta.Feature)
		assert.Equal(cfg.T(), "Publish course", cfg.Meta.Story)
		assert.Equal(cfg.T(), axiom.SeverityNormal, cfg.Meta.Severity)
		expectedTitle := "Go API Autotests"

		course := axiomCourse{
			Title:     expectedTitle,
			Published: false,
		}

		course.Published = true

		assert.True(cfg.T(), course.Published)
		assert.Equal(
			cfg.T(),
			[]string{"axiom", "courses", "publish", "regression"},
			cfg.Meta.Tags,
		)
		assert.Equal(cfg.T(), []string{"AXIOM-202"}, cfg.Meta.Issues)
		assert.Equal(cfg.T(), []string{"COURSES-002"}, cfg.Meta.TestCases)

		expectedLabels := map[string]string{
			"owner":     "courses-team",
			"component": "courses-service",
		}
		assert.Equal(cfg.T(), expectedLabels, cfg.Meta.Labels)
	})
}
