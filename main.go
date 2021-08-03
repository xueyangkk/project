package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"exams-api/configs"
	"exams-api/internal/router"
	"exams-api/pkg/env"
	"exams-api/pkg/logger"
	"exams-api/pkg/shutdown"

	"go.uber.org/zap"
)

// @title swagger 接口文档
// @version 2.0
// @description

// @contact.name
// @contact.url
// @contact.email

// @license.name MIT
// @license.url https://exams-api/blob/master/LICENSE

// @host 127.0.0.1:9999
// @BasePath
func main() {
	// 初始化 logger
	loggers, err := logger.NewJSONLogger(
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName(), env.Active().Value())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.ProjectLogFile()),
	)
	if err != nil {
		panic(err)
	}
	defer loggers.Sync()

	// 初始化 HTTP 服务
	s, err := router.NewHTTPServer(loggers)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.ProjectPort(),
		Handler: s.Mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			loggers.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				loggers.Error("server shutdown err", zap.Error(err))
			} else {
				loggers.Info("server shutdown success")
			}
		},

		// 关闭 db
		func() {
			if s.Db != nil {
				if err := s.Db.DbWClose(); err != nil {
					loggers.Error("dbw close err", zap.Error(err))
				} else {
					loggers.Info("dbw close success")
				}

				if err := s.Db.DbRClose(); err != nil {
					loggers.Error("dbr close err", zap.Error(err))
				} else {
					loggers.Info("dbr close success")
				}
			}
		},

		// 关闭 cache
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					loggers.Error("cache close err", zap.Error(err))
				} else {
					loggers.Info("cache close success")
				}
			}
		},

		// 关闭 gRPC client
		func() {
			if s.GrpClient != nil {
				if err := s.GrpClient.Conn().Close(); err != nil {
					loggers.Error("gRPC client close err", zap.Error(err))
				} else {
					loggers.Info("gRPC client close success")
				}
			}
		},
	)
}
