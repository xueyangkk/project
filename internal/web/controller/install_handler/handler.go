package install_handler

import (
	"exams-api/internal/pkg/core"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	View() core.HandlerFunc
	Execute() core.HandlerFunc
	Restart() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) i() {}
