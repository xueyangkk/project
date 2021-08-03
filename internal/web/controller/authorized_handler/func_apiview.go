package authorized_handler

import (
	"net/http"

	"exams-api/internal/api/code"
	"exams-api/internal/pkg/core"
	"exams-api/pkg/errno"
)

type apiViewRequest struct {
	Id string `uri:"id"` // 主键ID
}

type apiViewResponse struct {
	HashID string `json:"hash_id"` // hashID
}

func (h *handler) ApiView() core.HandlerFunc {
	return func(c core.Context) {
		req := new(apiViewRequest)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		obj := new(apiViewResponse)
		obj.HashID = req.Id

		c.HTML("authorized_api", obj)
	}
}
