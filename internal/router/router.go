package router

import (
	"exams-api/configs"
	"exams-api/internal/pkg/cache"
	"exams-api/internal/pkg/core"
	"exams-api/internal/pkg/db"
	"exams-api/internal/pkg/grpc"
	"exams-api/internal/pkg/metrics"
	"exams-api/internal/pkg/notify"
	"exams-api/internal/router/middleware"
	"exams-api/pkg/file"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type resource struct {
	mux     core.Mux
	logger  *zap.Logger
	db      db.Repo
	cache   cache.Repo
	grpConn grpc.ClientConn
	middles middleware.Middleware
}

type Server struct {
	Mux       core.Mux
	Db        db.Repo
	Cache     cache.Repo
	GrpClient grpc.ClientConn
}

func NewHTTPServer(logger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	r := new(resource)
	r.logger = logger

	openBrowserUri := "http://127.0.0.1" + configs.ProjectPort()

	_, ok := file.IsExists(configs.ProjectInstallFile())
	if !ok { // 未安装
		openBrowserUri += "/install"
	} else { // 已安装
		// 初始化 DB
		dbRepo, err := db.New()
		if err != nil {
			logger.Fatal("new db err", zap.Error(err))
		}
		r.db = dbRepo

		// 初始化 Cache
		cacheRepo, err := cache.New()
		if err != nil {
			logger.Fatal("new cache err", zap.Error(err))
		}
		r.cache = cacheRepo

		// 初始化 gRPC client
		gRPCRepo, err := grpc.New()
		if err != nil {
			logger.Fatal("new grpc err", zap.Error(err))
		}
		r.grpConn = gRPCRepo
	}

	mux, err := core.New(logger,
		core.WithEnableOpenBrowser(openBrowserUri),
		core.WithEnableCors(),
		core.WithEnableRate(),
		core.WithPanicNotify(notify.OnPanicNotify),
		core.WithRecordMetrics(metrics.RecordMetrics),
	)

	if err != nil {
		panic(err)
	}

	r.mux = mux
	r.middles = middleware.New(logger, r.cache, r.db)

	// 设置 WEB 路由
	setWebRouter(r)

	// 设置 API 路由
	setApiRouter(r)

	// 设置 GraphQL 路由
	setGraphQLRouter(r)

	s := new(Server)
	s.Mux = mux
	s.Db = r.db
	s.Cache = r.cache
	s.GrpClient = r.grpConn

	return s, nil
}
