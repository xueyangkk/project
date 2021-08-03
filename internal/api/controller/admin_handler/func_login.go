package admin_handler

import (
	"encoding/json"
	"net/http"
	"time"

	"exams-api/internal/api/code"
	"exams-api/internal/api/service/admin_service"
	"exams-api/internal/pkg/cache"
	"exams-api/internal/pkg/core"
	"exams-api/internal/pkg/password"
	"exams-api/pkg/errno"

	"github.com/pkg/errors"
)

type loginRequest struct {
	Username string `form:"username"` // 用户名
	Password string `form:"password"` // 密码
}

type loginResponse struct {
	Token string `json:"token"` // 用户身份标识
}

// Login 管理员登录
// @Summary 管理员登录
// @Description 管理员登录
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} loginResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/login [post]
func (h *handler) Login() core.HandlerFunc {
	return func(c core.Context) {
		req := new(loginRequest)
		res := new(loginResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		searchOneData := new(admin_service.SearchOneData)
		searchOneData.Username = req.Username
		searchOneData.Password = password.GeneratePassword(req.Password)
		searchOneData.IsUsed = 1

		info, err := h.adminService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(err),
			)
			return
		}

		if info == nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(errors.New("未查询出符合条件的用户")),
			)
			return
		}

		token := password.GenerateLoginToken(info.Id)

		// 用户信息
		adminJsonInfo, _ := json.Marshal(info)

		// 记录 Redis 中
		err = h.cache.Set(h.adminService.CacheKeyPrefix()+token, string(adminJsonInfo), time.Hour*24, cache.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithErr(err),
			)
			return
		}

		res.Token = token
		c.Payload(res)
	}
}
