package builders

import (
	coursev1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/et891/go-api-autotests/models"
)

type CourseCreate struct {
	Title           string
	MaxScore        int
	MinScore        int
	Description     string
	EstimatedTime   string
	PreviewFileID   string
	CreatedByUserID string
}

type CourseCreateOption func(*CourseCreate)

func (b *Builder) CourseCreate(options ...CourseCreateOption) CourseCreate {
	course := CourseCreate{
		Title:         b.generator.Title(),
		MaxScore:      b.generator.MaxScore(),
		MinScore:      b.generator.MinScore(),
		Description:   b.generator.Description(),
		EstimatedTime: b.generator.EstimatedTime(),
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
	return func(course *CourseCreate) { course.MaxScore = maxScore }
}

func WithCourseCreateMinScore(minScore int) CourseCreateOption {
	return func(course *CourseCreate) { course.MinScore = minScore }
}

func WithCourseCreateDescription(description string) CourseCreateOption {
	return func(course *CourseCreate) { course.Description = description }
}

func WithCourseCreateEstimatedTime(estimatedTime string) CourseCreateOption {
	return func(course *CourseCreate) { course.EstimatedTime = estimatedTime }
}

func WithCourseCreatePreviewFileID(fileID string) CourseCreateOption {
	return func(course *CourseCreate) { course.PreviewFileID = fileID }
}

func WithCourseCreateCreatedByUserID(userID string) CourseCreateOption {
	return func(course *CourseCreate) { course.CreatedByUserID = userID }
}

func (c CourseCreate) HTTPRequest() models.CreateCourseRequest {
	return models.CreateCourseRequest{
		Title:           c.Title,
		MaxScore:        &c.MaxScore,
		MinScore:        &c.MinScore,
		Description:     c.Description,
		EstimatedTime:   &c.EstimatedTime,
		PreviewFileID:   c.PreviewFileID,
		CreatedByUserID: c.CreatedByUserID,
	}
}

func (c CourseCreate) GRPCRequest() *coursev1.CreateCourseRequest {
	maxScore := int32(c.MaxScore)
	minScore := int32(c.MinScore)

	return &coursev1.CreateCourseRequest{
		Title:           c.Title,
		MaxScore:        &maxScore,
		MinScore:        &minScore,
		Description:     c.Description,
		EstimatedTime:   &c.EstimatedTime,
		PreviewFileId:   c.PreviewFileID,
		CreatedByUserId: c.CreatedByUserID,
	}
}

type CourseUpdate struct {
	Title         string
	MaxScore      int
	MinScore      int
	Description   string
	EstimatedTime string
}

type CourseUpdateOption func(*CourseUpdate)

func (b *Builder) CourseUpdate(options ...CourseUpdateOption) CourseUpdate {
	update := CourseUpdate{
		Title:         b.generator.Title(),
		MaxScore:      b.generator.MaxScore(),
		MinScore:      b.generator.MinScore(),
		Description:   b.generator.Description(),
		EstimatedTime: b.generator.EstimatedTime(),
	}

	for _, option := range options {
		option(&update)
	}

	return update
}

func WithCourseUpdateTitle(title string) CourseUpdateOption {
	return func(update *CourseUpdate) { update.Title = title }
}

func WithCourseUpdateMaxScore(maxScore int) CourseUpdateOption {
	return func(update *CourseUpdate) { update.MaxScore = maxScore }
}

func WithCourseUpdateMinScore(minScore int) CourseUpdateOption {
	return func(update *CourseUpdate) { update.MinScore = minScore }
}

func WithCourseUpdateDescription(description string) CourseUpdateOption {
	return func(update *CourseUpdate) { update.Description = description }
}

func WithCourseUpdateEstimatedTime(estimatedTime string) CourseUpdateOption {
	return func(update *CourseUpdate) { update.EstimatedTime = estimatedTime }
}

func (u CourseUpdate) HTTPRequest() models.UpdateCourseRequest {
	return models.UpdateCourseRequest{
		Title:         &u.Title,
		MaxScore:      &u.MaxScore,
		MinScore:      &u.MinScore,
		Description:   &u.Description,
		EstimatedTime: &u.EstimatedTime,
	}
}

func (u CourseUpdate) GRPCRequest(courseID string) *coursev1.UpdateCourseRequest {
	maxScore := int32(u.MaxScore)
	minScore := int32(u.MinScore)

	return &coursev1.UpdateCourseRequest{
		Id:            courseID,
		Title:         &u.Title,
		MaxScore:      &maxScore,
		MinScore:      &minScore,
		Description:   &u.Description,
		EstimatedTime: &u.EstimatedTime,
	}
}
