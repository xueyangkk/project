package tool_handler

import (
	"exams-api/configs"
	"exams-api/internal/pkg/core"
)

func (h *handler) HashIdsView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("tool_hashids", configs.Get())
	}
}
