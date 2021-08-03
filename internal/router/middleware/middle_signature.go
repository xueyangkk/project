package middleware

import (
	"net/http"
	"strings"
	"time"

	"exams-api/internal/api/code"
	"exams-api/internal/pkg/core"
	"exams-api/pkg/errno"
	"exams-api/pkg/signature"

	"github.com/koketama/urltable"
	"github.com/pkg/errors"
)

const ttl = time.Minute * 2 // 签名超时时间 2 分钟

var whiteListPath = map[string]bool{
	"/login/web": true,
}

func (m *middleware) Signature() core.HandlerFunc {
	return func(c core.Context) {
		// 签名信息
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New("Header 中缺少 Authorization 参数")),
			)
			return
		}

		// 时间信息
		date := c.GetHeader("Authorization-Date")
		if date == "" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New("Header 中缺少 Date 参数")),
			)
			return
		}

		// 通过签名信息获取 key
		authorizationSplit := strings.Split(authorization, " ")
		if len(authorizationSplit) < 2 {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New("Header 中 Authorization 格式错误")),
			)
			return
		}

		key := authorizationSplit[0]

		data, err := m.authorizedService.DetailByKey(c, key)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(err),
			)
			return
		}

		// 验证 cache 是否被调用
		if data.IsUsed == -1 {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New(key + " 已被禁止调用")),
			)
			return
		}

		if len(data.Apis) < 1 {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New(key + " 未进行接口授权")),
			)
			return
		}

		if !whiteListPath[c.Path()] {
			// 验证 c.Method() + c.Path() 是否授权
			table := urltable.NewTable()
			for _, v := range data.Apis {
				_ = table.Append(v.Method + v.Api)
			}

			if pattern, _ := table.Mapping(c.Method() + c.Path()); pattern == "" {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.SignatureError,
					code.Text(code.SignatureError)).WithErr(errors.New(c.Method() + c.Path() + " 未进行接口授权")),
				)
				return
			}
		}

		ok, err := signature.New(key, data.Secret, ttl).Verify(authorization, date, c.Path(), c.Method(), c.RequestInputParams())
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(err),
			)
			return
		}

		if !ok {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.SignatureError,
				code.Text(code.SignatureError)).WithErr(errors.New("Header 中 Authorization 信息错误")),
			)
			return
		}
	}
}
