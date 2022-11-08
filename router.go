package main

import (
	"github.com/jamesluo111/core_web/framework/gin"
	"github.com/jamesluo111/core_web/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	//http+静态路由
	core.GET("/user/login", middleware.Test3(), UserLoginController)
	core.GET("/user/id", middleware.Test3(), GetUserController)
	//批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

	}
}
