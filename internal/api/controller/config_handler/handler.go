package config_handler

import (
	"exams-api/internal/pkg/cache"
	"exams-api/internal/pkg/core"
	"exams-api/internal/pkg/db"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Email 修改邮件配置
	// @Tags API.config
	// @Router /api/config/email [patch]
	Email() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	cache  cache.Repo
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger: logger,
		cache:  cache,
	}
}

func (h *handler) i() {}
