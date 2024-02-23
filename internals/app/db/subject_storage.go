package db

import (
	"context"
	"courses_serve/example/internals/app/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type SubjectSrorage struct {
	databasePool *pgxpool.Pool
}

func NewSubjectsStorage(pool *pgxpool.Pool) *SubjectSrorage {
	storage := new(SubjectSrorage)
	storage.databasePool = pool
	return storage
}

func (storage *SubjectSrorage) GetAllSubjects() []models.Subject {
	query := "SELECT id, name, description FROM subjects"

	var result []models.Subject
	err := pgxscan.Select(context.Background(), storage.databasePool, &result, query)
	if err != nil {
		log.Errorln(err)
	}
	return result

}
