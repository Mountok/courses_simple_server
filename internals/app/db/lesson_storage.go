package db

import (
	"context"
	"courses_serve/example/internals/app/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type LessonsStorage struct {
	databasePool *pgxpool.Pool
}

func NewLessonsStorage(pool *pgxpool.Pool) *LessonsStorage {
	storage := new(LessonsStorage)
	storage.databasePool = pool
	return storage
}

func (storage *LessonsStorage) GetLessonsById(subjectId int) []models.Lesson {
	query := "select id, name, description, positions from lessons where subject_id = $1 order by positions;"
	var result []models.Lesson
	err := pgxscan.Select(context.Background(), storage.databasePool, &result, query, subjectId)
	if err != nil {
		log.Errorln(err)
	}
	return result
}


func (storage *LessonsStorage) GetLessonContent(subjectId, lessonId int) []models.LessonContent {
	query := "select id, image, content, subject_id, lesson_id from lessons_contents where subject_id = $1 and lesson_id = $2"
	var result []models.LessonContent
	err := pgxscan.Select(context.Background(),storage.databasePool,&result,query,subjectId,lessonId)
	if err != nil {
		log.Errorln(err)
	}
	return result
}