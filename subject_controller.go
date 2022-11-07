package main

import (
	"fmt"
	"github.com/jamesluo111/core_web/framework/gin"
)

func SubjectListController(ctx *gin.Context) error {
	ctx.ISetOkStatus().IJson("ok")
	return nil
}

func SubjectDelController(ctx *gin.Context) error {
	ctx.ISetOkStatus().IJson("ok")
	return nil
}

func SubjectUpdateController(ctx *gin.Context) error {
	ctx.ISetOkStatus().IJson("ok")
	return nil
}

func SubjectGetController(ctx *gin.Context) error {
	id, _ := ctx.DefaultParamInt64("id", 0)
	fmt.Println(id)
	ctx.ISetOkStatus().IJson("ok")
	return nil
}
