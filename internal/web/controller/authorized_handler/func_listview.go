package authorized_handler

import "exams-api/internal/pkg/core"

func (h *handler) ListView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("authorized_list", nil)
	}
}
