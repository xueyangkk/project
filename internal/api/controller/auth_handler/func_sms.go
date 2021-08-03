package auth_handler

import (
	"exams-api/internal/pkg/core"
)



// SendAuthCode 发送验证码
// @Summary 发送验证码
// @Description 发送验证码
// @Tags API.admin
// @Accept json
// @Produce json
// @Param page query int false "第几页"
// @Param page_size query string false "每页显示条数"
// @Param username query string false "用户名"
// @Param nickname query string false "昵称"
// @Param mobile query string false "手机号"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [get]
func (h *handler) SendAuthCode() core.HandlerFunc {

}
