package user_handler

import (
	"exams-api/internal/pkg/cache"
	"exams-api/internal/pkg/db"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	// i 为了避免被其他包实现
	i()

	//// Create 创建用户
	//// @Tags Test
	//// @Router /test/create [post]
	//Create() core.HandlerFunc
	//
	//// Update 编辑用户
	//// @Tags Test
	//// @Router /test/update [post]
	//Update() core.HandlerFunc
	//
	//// Delete 删除用户
	//// @Tags Test
	//// @Router /test/delete [post]
	//Delete() core.HandlerFunc
	//
	//// Detail 用户详情
	//// @Tags Test
	//// @Router /test/detail [post]
	//Detail() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       cache.Repo
	//userService user_service.UserService
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		//userService: user_service.NewUserService(db, cache),
	}
}

func (h *handler) i() {}