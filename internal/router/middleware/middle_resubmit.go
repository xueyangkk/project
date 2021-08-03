package middleware

import (
	"net/http"
	"time"

	"exams-api/configs"
	"exams-api/internal/api/code"
	"exams-api/internal/pkg/cache"
	"exams-api/internal/pkg/core"
	"exams-api/pkg/errno"
	"exams-api/pkg/token"

	"github.com/pkg/errors"
)

func (m *middleware) Resubmit() core.HandlerFunc {

	redisKeyPrefix := configs.ProjectName() + ":request-id:"

	return func(c core.Context) {
		cfg := configs.Get().URLToken

		tokenString, err := token.New(cfg.Secret).UrlSign(c.Path(), c.Method(), c.RequestInputParams())
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ResubmitError,
				code.Text(code.ResubmitError)).WithErr(err),
			)
			return
		}

		redisKey := redisKeyPrefix + tokenString
		if !m.cache.Exists(redisKey) {
			err = m.cache.Set(redisKey, "1", time.Minute*cfg.ExpireDuration)
			if err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.ResubmitError,
					code.Text(code.ResubmitError)).WithErr(err),
				)
				return
			}

			return
		}

		redisValue, err := m.cache.Get(redisKey, cache.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ResubmitError,
				code.Text(code.ResubmitError)).WithErr(err),
			)
			return
		}

		if redisValue == "1" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ResubmitMsg,
				code.Text(code.ResubmitMsg)).WithErr(errors.New("resubmit")),
			)
			return
		}

		return
	}
}
