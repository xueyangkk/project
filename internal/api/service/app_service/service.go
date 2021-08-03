package app_service

import (
	"exams-api/internal/pkg/cache"
	"exams-api/internal/pkg/db"
)

type Service interface {
	i()
}

type service struct {
	db    db.Repo
	cache cache.Repo
}

func New(db db.Repo, cache cache.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {

}
