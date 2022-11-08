package main

import (
	"fmt"
	"github.com/jamesluo111/core_web/framework/gin"
	"github.com/jamesluo111/core_web/provider/demo"
)

func SubjectListController(ctx *gin.Context) {
	//获取demo实例
	demoService := ctx.MustMake(demo.Key).(demo.Service)

	//调用服务实例的方法
	foo := demoService.GetFoo()
	//输出结果
	ctx.ISetOkStatus().IJson(foo)
	return
}

func SubjectDelController(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("ok")
	return
}

func SubjectUpdateController(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("ok")
	return
}

func SubjectGetController(ctx *gin.Context) {
	id, _ := ctx.DefaultParamInt64("id", 0)
	fmt.Println(id)
	ctx.ISetOkStatus().IJson("ok")
	return
}
