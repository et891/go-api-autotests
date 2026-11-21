package builders

import (
	coursev1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/et891/go-api-autotests/models"
)

func intToInt32Ptr(value *int) *int32 {
	if value == nil {
		return nil
	}

	converted := int32(*value)
	return &converted
}

type CourseCreate struct {
	Title           string
	MaxScore        *int
	MinScore        *int
	Description     string
	EstimatedTime   *string
	PreviewFileID   string
	CreatedByUserID string
}

type CourseCreateOption func(*CourseCreate)

func (b *Builder) CourseCreate(previewFileID string, createdByUserID string, options ...CourseCreateOption) CourseCreate {
	title := b.generator.Title()
	maxScore := b.generator.MaxScore()
	minScore := b.generator.MinScore()
	description := b.generator.Description()
	estimatedTime := b.generator.EstimatedTime()

	course := CourseCreate{
		Title:           title,
		MaxScore:        &maxScore,
		MinScore:        &minScore,
		Description:     description,
		EstimatedTime:   &estimatedTime,
		PreviewFileID:   previewFileID,
		CreatedByUserID: createdByUserID,
	}

	for _, option := range options {
		option(&course)
	}

	return course
}

func WithCourseCreateTitle(title string) CourseCreateOption {
	return func(course *CourseCreate) { course.Title = title }
}

func WithCourseCreateMaxScore(maxScore int) CourseCreateOption {
	return func(course *CourseCreate) { course.MaxScore = &maxScore }
}

func WithCourseCreateMinScore(minScore int) CourseCreateOption {
	return func(course *CourseCreate) { course.MinScore = &minScore }
}

func WithCourseCreateDescription(description string) CourseCreateOption {
	return func(course *CourseCreate) { course.Description = description }
}

func WithCourseCreateEstimatedTime(estimatedTime string) CourseCreateOption {
	return func(course *CourseCreate) { course.EstimatedTime = &estimatedTime }
}

func (c CourseCreate) HTTPRequest() models.CreateCourseRequest {
	return models.CreateCourseRequest{
		Title:           c.Title,
		MaxScore:        c.MaxScore,
		MinScore:        c.MinScore,
		Description:     c.Description,
		EstimatedTime:   c.EstimatedTime,
		PreviewFileID:   c.PreviewFileID,
		CreatedByUserID: c.CreatedByUserID,
	}
}

func (c CourseCreate) GRPCRequest() *coursev1.CreateCourseRequest {
	return &coursev1.CreateCourseRequest{
		Title:           c.Title,
		MaxScore:        intToInt32Ptr(c.MaxScore),
		MinScore:        intToInt32Ptr(c.MinScore),
		Description:     c.Description,
		EstimatedTime:   c.EstimatedTime,
		PreviewFileId:   c.PreviewFileID,
		CreatedByUserId: c.CreatedByUserID,
	}
}

type CourseUpdate struct {
	Title         *string
	MaxScore      *int
	MinScore      *int
	Description   *string
	EstimatedTime *string
}

type CourseUpdateOption func(*CourseUpdate)

func (b *Builder) CourseUpdate(options ...CourseUpdateOption) CourseUpdate {
	title := b.generator.Title()
	maxScore := b.generator.MaxScore()
	minScore := b.generator.MinScore()
	description := b.generator.Description()
	estimatedTime := b.generator.EstimatedTime()

	update := CourseUpdate{
		Title:         &title,
		MaxScore:      &maxScore,
		MinScore:      &minScore,
		Description:   &description,
		EstimatedTime: &estimatedTime,
	}

	for _, option := range options {
		option(&update)
	}

	return update
}

func WithCourseUpdateTitle(title string) CourseUpdateOption {
	return func(course *CourseUpdate) { course.Title = &title }
}

func WithCourseUpdateMaxScore(maxScore int) CourseUpdateOption {
	return func(course *CourseUpdate) { course.MaxScore = &maxScore }
}
func WithCourseUpdateMinScore(minScore int) CourseUpdateOption {
	return func(course *CourseUpdate) { course.MinScore = &minScore }
}
func WithCourseUpdateDescription(description string) CourseUpdateOption {
	return func(course *CourseUpdate) { course.Description = &description }
}
func WithCourseUpdateEstimatedTime(estimatedTime string) CourseUpdateOption {
	return func(course *CourseUpdate) { course.EstimatedTime = &estimatedTime }
}
func (c CourseUpdate) HTTPRequest() models.UpdateCourseRequest {
	return models.UpdateCourseRequest{
		Title:         c.Title,
		MaxScore:      c.MaxScore,
		MinScore:      c.MinScore,
		Description:   c.Description,
		EstimatedTime: c.EstimatedTime,
	}
}

func (c CourseUpdate) GRPCRequest(courseID string) *coursev1.UpdateCourseRequest {
	return &coursev1.UpdateCourseRequest{
		Id:            courseID,
		Title:         c.Title,
		MaxScore:      intToInt32Ptr(c.MaxScore),
		MinScore:      intToInt32Ptr(c.MinScore),
		Description:   c.Description,
		EstimatedTime: c.EstimatedTime,
	}
}
