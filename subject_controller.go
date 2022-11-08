package main

import (
	"fmt"
	"github.com/jamesluo111/core_web/framework/gin"
)

func SubjectListController(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("ok")
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
