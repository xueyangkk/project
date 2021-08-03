package middleware

import (
	"encoding/json"
	"net/http"

	"exams-api/internal/api/code"
	"exams-api/internal/pkg/cache"
	"exams-api/internal/pkg/core"
	"exams-api/pkg/errno"

	"github.com/pkg/errors"
)

func (m *middleware) Token(ctx core.Context) (userId int64, userName string, err errno.Error) {
	token := ctx.GetHeader("Token")
	if token == "" {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("Header 中缺少 Token 参数"))

		return
	}

	if !m.cache.Exists(m.adminService.CacheKeyPrefix() + token) {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("请先登录"))

		return
	}

	cacheData, cacheErr := m.cache.Get(m.adminService.CacheKeyPrefix()+token, cache.WithTrace(ctx.Trace()))
	if cacheErr != nil {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(cacheErr)

		return
	}

	type userInfo struct {
		Id       int64  `json:"id"`       // 用户ID
		Username string `json:"username"` // 用户名
	}

	var userData userInfo
	_ = json.Unmarshal([]byte(cacheData), &userData)

	userId = userData.Id
	userName = userData.Username

	return
}
