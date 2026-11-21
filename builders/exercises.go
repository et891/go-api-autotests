package builders

import (
	"strconv"

	coursev1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/et891/go-api-autotests/models"
)

type ExerciseCreate struct {
	Title         string
	CourseId      int
	MaxScore      *int
	MinScore      *int
	OrderIndex    int
	Description   string
	EstimatedTime *string
}

type ExerciseCreateOption func(*ExerciseCreate)

func (b *Builder) ExerciseCreate(courseID int, options ...ExerciseCreateOption) ExerciseCreate {
	title := b.generator.Title()
	maxScore := b.generator.MaxScore()
	minScore := b.generator.MinScore()
	orderIndex := b.generator.OrderIndex()
	description := b.generator.Description()
	estimatedTime := b.generator.EstimatedTime()

	exercise := ExerciseCreate{
		Title:         title,
		CourseId:      courseID,
		MaxScore:      &maxScore,
		MinScore:      &minScore,
		OrderIndex:    orderIndex,
		Description:   description,
		EstimatedTime: &estimatedTime,
	}

	for _, option := range options {
		option(&exercise)
	}

	return exercise
}

func WithExerciseCreateTitle(title string) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.Title = title }
}

func WithExerciseCreateMaxScore(maxScore int) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.MaxScore = &maxScore }
}

func WithExerciseCreateMinScore(minScore int) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.MinScore = &minScore }
}

func WithExerciseCreateOrderIndex(orderIndex int) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.OrderIndex = orderIndex }
}

func WithExerciseCreateDescription(description string) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.Description = description }
}

func WithExerciseCreateEstimatedTime(estimatedTime string) ExerciseCreateOption {
	return func(exercise *ExerciseCreate) { exercise.EstimatedTime = &estimatedTime }
}

func (c ExerciseCreate) HTTPRequest() models.CreateExerciseRequest {
	return models.CreateExerciseRequest{
		Title:         c.Title,
		CourseID:      c.CourseId,
		MaxScore:      c.MaxScore,
		MinScore:      c.MinScore,
		OrderIndex:    c.OrderIndex,
		Description:   c.Description,
		EstimatedTime: c.EstimatedTime,
	}
}

func (c ExerciseCreate) GRPCRequest() *coursev1.CreateExerciseRequest {
	return &coursev1.CreateExerciseRequest{
		Title:         c.Title,
		CourseId:      strconv.Itoa(c.CourseId),
		MaxScore:      intToInt32Ptr(c.MaxScore),
		MinScore:      intToInt32Ptr(c.MinScore),
		OrderIndex:    int32(c.OrderIndex),
		Description:   c.Description,
		EstimatedTime: c.EstimatedTime,
	}
}

type ExerciseUpdate struct {
	Title         *string
	MaxScore      *int
	MinScore      *int
	OrderIndex    *int
	Description   *string
	EstimatedTime *string
}

type ExerciseUpdateOption func(*ExerciseUpdate)

func (b *Builder) ExerciseUpdate(options ...ExerciseUpdateOption) ExerciseUpdate {
	title := b.generator.Title()
	maxScore := b.generator.MaxScore()
	minScore := b.generator.MinScore()
	orderIndex := b.generator.OrderIndex()
	description := b.generator.Description()
	estimatedTime := b.generator.EstimatedTime()

	update := ExerciseUpdate{
		Title:         &title,
		MaxScore:      &maxScore,
		MinScore:      &minScore,
		OrderIndex:    &orderIndex,
		Description:   &description,
		EstimatedTime: &estimatedTime,
	}

	for _, option := range options {
		option(&update)
	}

	return update
}

func WithExerciseUpdateTitle(title string) ExerciseUpdateOption {
	return func(exercise *ExerciseUpdate) { exercise.Title = &title }
}

func WithExerciseUpdateMaxScore(maxScore int) ExerciseUpdateOption {
	return func(exercise *ExerciseUpdate) { exercise.MaxScore = &maxScore }
}
func WithExerciseUpdateMinScore(minScore int) ExerciseUpdateOption {
	return func(exercise *ExerciseUpdate) { exercise.MinScore = &minScore }
}
func WithExerciseUpdateOrderIndex(orderIndex int) ExerciseUpdateOption {
	return func(exercise *ExerciseUpdate) { exercise.OrderIndex = &orderIndex }
}
func WithExerciseUpdateDescription(description string) ExerciseUpdateOption {
	return func(exercise *ExerciseUpdate) { exercise.Description = &description }
}
func WithExerciseUpdateEstimatedTime(estimatedTime string) ExerciseUpdateOption {
	return func(exercise *ExerciseUpdate) { exercise.EstimatedTime = &estimatedTime }
}
func (c ExerciseUpdate) HTTPRequest() models.UpdateExerciseRequest {
	return models.UpdateExerciseRequest{
		Title:         c.Title,
		MaxScore:      c.MaxScore,
		MinScore:      c.MinScore,
		Description:   c.Description,
		OrderIndex:    c.OrderIndex,
		EstimatedTime: c.EstimatedTime,
	}
}

func (c ExerciseUpdate) GRPCRequest(exerciseID string) *coursev1.UpdateExerciseRequest {
	return &coursev1.UpdateExerciseRequest{
		Id:            exerciseID,
		Title:         c.Title,
		MaxScore:      intToInt32Ptr(c.MaxScore),
		MinScore:      intToInt32Ptr(c.MinScore),
		OrderIndex:    intToInt32Ptr(c.OrderIndex),
		Description:   c.Description,
		EstimatedTime: c.EstimatedTime,
	}
}
