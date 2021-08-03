package auth_handler

import (
	"exams-api/configs"
	"exams-api/internal/pkg/cache"
	"exams-api/internal/pkg/core"
	"exams-api/internal/pkg/db"
	"exams-api/pkg/hash"

	"go.uber.org/zap"
)



type Handler interface {

	// i 为了避免被其他包实现
	i()

	LoginWithAuthCode() core.HandlerFunc
	SendAuthCode() core.HandlerFunc
	
	//LoginWithPassword() core.HandlerFunc
	//LoginWithUnionId() core.HandlerFunc
	//
	//SignUp() core.HandlerFunc



}

type handler struct {
	logger       *zap.Logger
	cache        cache.Repo
	hashids      hash.Hash
	//adminService admin_service.Service
}




func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {

	return &handler{
		logger:       logger,
		cache:        cache,
		hashids:      hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		//adminService: admin_service.New(db, cache),
	}

}

func (h *handler) i() {}