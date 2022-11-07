package main

import (
	"github.com/jamesluo111/core_web/framework"
	"github.com/jamesluo111/core_web/framework/middleware"
)

func registerRouter(core *framework.Core) {
	//http+静态路由
	core.Get("/user/login", middleware.Test3(), UserLoginController)
	core.Get("/user/id", middleware.Test3(), GetUserController)
	//批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

	}
}
