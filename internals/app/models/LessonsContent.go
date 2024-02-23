package models


type LessonContent struct {
	Id int `json:"id"`
	Image string `json:"image"`
	Content string `json:"content"`
	SubjectId int `json:"subject_id"`
	LessonId int `json:"lesson_id"`
}