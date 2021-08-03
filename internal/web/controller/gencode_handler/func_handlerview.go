package gencode_handler

import (
	"exams-api/internal/pkg/core"
)

func (h *handler) HandlerView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("gencode_handler", nil)
	}
}
