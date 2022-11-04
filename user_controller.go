package main

import (
	"jamesluo1/framework"
)

func UserLoginController(ctx *framework.Context) error {
	ctx.Json(200, "登陆成功")
	return nil
}
