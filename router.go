package main

import "jamesluo1/framework"

func registerRouter(core *framework.Core) {
	//http+静态路由
	core.Get("/user/login", UserLoginController)
	//批量通用前缀
	sujectApi := core.Group("subject")
	{
		sujectApi.Get("/list", SubjectListController)
	}
}
