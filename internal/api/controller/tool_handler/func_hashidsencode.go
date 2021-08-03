package tool_handler

import (
	"net/http"

	"exams-api/internal/api/code"
	"exams-api/internal/pkg/core"
	"exams-api/pkg/errno"

	"github.com/spf13/cast"
)

type hashIdsEncodeRequest struct {
	Id int32 `uri:"id"` // 需加密的数字
}

type hashIdsEncodeResponse struct {
	Val string `json:"val"` // 加密后的值
}

// HashIdsEncode HashIds 加密
// @Summary HashIds 加密
// @Description HashIds 加密
// @Tags API.tool
// @Accept json
// @Produce json
// @Param id path string true "需加密的数字"
// @Success 200 {object} hashIdsEncodeResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/hashids/encode/{id} [get]
func (h *handler) HashIdsEncode() core.HandlerFunc {
	return func(c core.Context) {
		req := new(hashIdsEncodeRequest)
		res := new(hashIdsEncodeResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(req.Id)})
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		res.Val = hashId

		c.Payload(res)
	}
}
