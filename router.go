package main

import "jamesluo1/framework"

func registerRouter(core *framework.Core) {
	//http+静态路由
	core.Get("/user/login", UserLoginController)
	//批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

	}
}
