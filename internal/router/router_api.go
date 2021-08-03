package router

import (
	"exams-api/internal/api/controller/admin_handler"
	"exams-api/internal/api/controller/auth_handler"
	"exams-api/internal/api/controller/authorized_handler"
	"exams-api/internal/api/controller/config_handler"
	"exams-api/internal/api/controller/tool_handler"
	"exams-api/internal/pkg/core"
)

func setApiRouter(r *resource) {

	// admin
	adminHandler := admin_handler.New(r.logger, r.db, r.cache)

	// login
	login := r.mux.Group("/login", r.middles.Signature())
	{
		login.POST("/web", adminHandler.Login())

	}

	// api
	api := r.mux.Group("/api", core.WrapAuthHandler(r.middles.Token), r.middles.Signature())
	{
		// authorized
		authorizedHandler := authorized_handler.New(r.logger, r.db, r.cache)
		api.POST("/authorized", authorizedHandler.Create())
		api.GET("/authorized", authorizedHandler.List())
		api.PATCH("/authorized/used", authorizedHandler.UpdateUsed())
		api.DELETE("/authorized/:id", core.AliasForRecordMetrics("/api/authorized/info"), authorizedHandler.Delete())

		api.POST("/authorized_api", authorizedHandler.CreateAPI())
		api.GET("/authorized_api", authorizedHandler.ListAPI())
		api.DELETE("/authorized_api/:id", core.AliasForRecordMetrics("/api/authorized_api/info"), authorizedHandler.DeleteAPI())

		api.POST("/admin", adminHandler.Create())
		api.GET("/admin", adminHandler.List())
		api.PATCH("/admin/used", adminHandler.UpdateUsed())
		api.PATCH("/admin/reset_password/:id", adminHandler.ResetPassword())
		api.DELETE("/admin/:id", adminHandler.Delete())
		api.POST("/admin/logout", adminHandler.Logout())
		api.PATCH("/admin/modify_password", adminHandler.ModifyPassword())
		api.GET("/admin/info", adminHandler.Detail())
		api.PATCH("/admin/modify_personal_info", adminHandler.ModifyPersonalInfo())

		// tool
		toolHandler := tool_handler.New(r.logger, r.db, r.cache)
		api.GET("/tool/hashids/encode/:id", toolHandler.HashIdsEncode())
		api.GET("/tool/hashids/decode/:id", toolHandler.HashIdsDecode())

		// config
		configHandler := config_handler.New(r.logger, r.db, r.cache)
		api.PATCH("/config/email", configHandler.Email())

	}

	authHandler := auth_handler.New(r.logger,r.db,r.cache)

	//登录接口
	authLogin  := r.mux.Group("/auth")
	{
		//send message or login
		authLogin.POST("/send/code",authHandler.SendAuthCode())
		authLogin.POST("/login/auth/code,",authHandler.LoginWithAuthCode())


	}


	 r.mux.Group("/v1",core.WrapAuthHandler(r.middles.Token))
	{



	}





}
