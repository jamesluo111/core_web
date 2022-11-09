package http

import (
	"github.com/jamesluo111/core_web/app/http/module/demo"
	"github.com/jamesluo111/core_web/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	demo.Register(r)
}
