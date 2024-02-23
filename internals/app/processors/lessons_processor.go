package processors

import (
	"courses_serve/example/internals/app/db"
	"courses_serve/example/internals/app/models"
	"errors"
)

type LesssonsProgessor struct {
	storage *db.LessonsStorage
}

func NewLessonsProcessor(storage *db.LessonsStorage) *LesssonsProgessor {
	processor := new(LesssonsProgessor)
	processor.storage = storage
	return processor
}

func (processor *LesssonsProgessor) ListLessons(subjectId int) ([]models.Lesson, error) {
	if subjectId > 0 {
		return processor.storage.GetLessonsById(subjectId),nil
	}
	return []models.Lesson{}, errors.New("subject id <= 0")
}

func (processor *LesssonsProgessor) LessonContent(subjectId,lessonId int) ([]models.LessonContent,error) {
	if subjectId > 0 && lessonId > 0 {
		return processor.storage.GetLessonContent(subjectId,lessonId), nil	
	}
	return []models.LessonContent{}, errors.New("not found")
}