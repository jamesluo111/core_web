package main

import (
	"fmt"
	"jamesluo1/framework"
)

func SubjectListController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("ok")
	return nil
}

func SubjectDelController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("ok")
	return nil
}

func SubjectUpdateController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("ok")
	return nil
}

func SubjectGetController(ctx *framework.Context) error {
	id, _ := ctx.ParamInt("id", 0)
	fmt.Println(id)
	ctx.SetOkStatus().Json("ok")
	return nil
}
