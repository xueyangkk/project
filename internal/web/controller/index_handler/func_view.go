package index_handler

import "exams-api/internal/pkg/core"

func (h *handler) View() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("index", nil)
	}
}
