package models

type Lesson struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Positions int `json:"positions"`
	SubjectsId int `json:"subject_id"`
}