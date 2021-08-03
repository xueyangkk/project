package admin_handler

import "exams-api/internal/pkg/core"

func (h *handler) AddView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("admin_add", nil)
	}
}
