package processors

import (
	"courses_serve/example/internals/app/db"
	"courses_serve/example/internals/app/models"
)

type SubjectsProcessor struct {
	storage *db.SubjectSrorage
}

func NewSubjectsProcessor(storage *db.SubjectSrorage) *SubjectsProcessor {
	processor := new(SubjectsProcessor)
	processor.storage = storage
	return processor
}

func (processor *SubjectsProcessor) ListSubjects() ([]models.Subject, error) {
	return processor.storage.GetAllSubjects(), nil
}
