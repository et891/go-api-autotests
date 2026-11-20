package tests

import (
	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
)

type createCourseParams struct {
	Title string
}

func (s *AxiomSuite) TestCourseCanBeCreatedWithDifferentEmails() {
	testCases := []axiom.Case{
		axiom.NewCase(
			axiom.WithCaseName("course can be created with Go API Autotests title"),
			axiom.WithCaseParams(createCourseParams{
				Title: "Go API Autotests",
			}),
			axiom.WithCaseContext(
				axiom.WithContextRaw(s.T().Context()),
				axiom.WithContextData("operation", "create-course"),
			),
		),
		axiom.NewCase(
			axiom.WithCaseName("course can be created with gRPC API Testing title"),
			axiom.WithCaseParams(createCourseParams{
				Title: "gRPC API Testing",
			}),
			axiom.WithCaseContext(
				axiom.WithContextRaw(s.T().Context()),
				axiom.WithContextData("operation", "create-course"),
			),
		),
		axiom.NewCase(
			axiom.WithCaseName("course can be created with Axiom Framework title"),
			axiom.WithCaseParams(createCourseParams{
				Title: "Axiom Framework",
			}),
			axiom.WithCaseContext(
				axiom.WithContextRaw(s.T().Context()),
				axiom.WithContextData("operation", "create-course"),
			),
		),
	}

	for _, testCase := range testCases {
		s.RunCase(testCase, coursesToolset.Action(func(cfg *axiom.Config, tools *axiomCoursesTools) {
			params := axiom.GetParams[createCourseParams](cfg)

			var course *axiomCourse

			cfg.Step("create course", func() {
				course = tools.Client.Create(axiomCourseData{
					Title: params.Title,
				})
			})

			cfg.Step("check created course", func() {
				assert.Equal(cfg.T(), params.Title, course.Title)
				assert.False(cfg.T(), course.Published)
			})

			cfg.Teardown("delete course", func() {
				tools.Client.Delete(course)
			})
		}))
	}
}
