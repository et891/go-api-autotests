package models

type Exercise struct {
	Id            int     `json:"id"`
	Title         string  `json:"title"`
	CourseId      int     `json:"course_id"`
	MaxScore      *int    `json:"maxScore,omitempty"`
	MinScore      *int    `json:"minScore,omitempty"`
	OrderIndex    int     `json:"orderIndex"`
	Description   string  `json:"description"`
	EstimatedTime *string `json:"estimatedTime,omitempty"`
}

type ExerciseResponse struct {
	Exercise Exercise `json:"exercise"`
}
type ExercisesResponse struct {
	Exercises []Exercise `json:"exercises"`
}
type ListExercisesQuery struct {
	CourseID string
}

func (q ListExercisesQuery) ToQueryParams() map[string]string {
	return map[string]string{"courseId": q.CourseID}
}

type CreateExerciseRequest struct {
	Title         string  `json:"title"`
	CourseID      int     `json:"courseId"`
	MaxScore      *int    `json:"maxScore,omitempty"`
	MinScore      *int    `json:"minScore,omitempty"`
	OrderIndex    int     `json:"orderIndex"`
	Description   string  `json:"description"`
	EstimatedTime *string `json:"estimatedTime,omitempty"`
}

type UpdateExerciseRequest struct {
	Id            int     `json:"id"`
	Title         *string `json:"title,omitempty"`
	MaxScore      *int    `json:"maxScore,omitempty"`
	MinScore      *int    `json:"minScore,omitempty"`
	OrderIndex    *int    `json:"orderIndex,omitempty"`
	Description   *string `json:"description,omitempty"`
	EstimatedTime *string `json:"estimatedTime,omitempty"`
}
