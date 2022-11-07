package main

import (
	"jamesluo1/framework"
	"time"
)

func UserLoginController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("登录成功")
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

func GetUserController(ctx *framework.Context) error {
	data := TodoPages{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	time.Sleep(10 * time.Second)
	ctx.Html("public/user.html", "user.html", data)
	ctx.SetOkStatus()
	return nil
}
