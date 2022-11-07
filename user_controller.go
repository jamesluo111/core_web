package main

import (
	"github.com/jamesluo111/core_web/framework/gin"
	"time"
)

func UserLoginController(ctx *gin.Context) error {
	ctx.ISetOkStatus().IJson("登录成功")
	return nil
}

type TodoPages struct {
	PageTitle string
	Todos     []Todo
}

type Todo struct {
	Title string
	Done  bool
}

func GetUserController(ctx *gin.Context) error {
	data := TodoPages{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	time.Sleep(10 * time.Second)
	ctx.IHtml("public/user.html", "user.html", data)
	ctx.ISetOkStatus()
	return nil
}
