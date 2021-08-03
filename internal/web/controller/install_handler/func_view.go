package install_handler

import (
	"exams-api/configs"
	"exams-api/internal/pkg/core"
)

func (h *handler) View() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("install_view", configs.Get())
	}
}
