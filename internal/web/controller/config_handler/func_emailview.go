package config_handler

import (
	"exams-api/configs"
	"exams-api/internal/pkg/core"
)

func (h *handler) EmailView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("config_email", configs.Get())
	}
}
