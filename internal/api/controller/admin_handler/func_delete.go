package admin_handler

import (
	"net/http"

	"exams-api/internal/api/code"
	"exams-api/internal/pkg/core"
	"exams-api/pkg/errno"
)

type deleteRequest struct {
	Id string `uri:"id"` // HashID
}

type deleteResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Delete 删除管理员
// @Summary 删除管理员
// @Description 删除管理员
// @Tags API.admin
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/{id} [delete]
func (h *handler) Delete() core.HandlerFunc {
	return func(c core.Context) {
		req := new(deleteRequest)
		res := new(deleteResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		id := int32(ids[0])

		err = h.adminService.Delete(c, id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminDeleteError,
				code.Text(code.AdminDeleteError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
