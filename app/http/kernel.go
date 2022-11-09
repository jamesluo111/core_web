package http

import "github.com/jamesluo111/core_web/framework/gin"

// NewHttpEngine 创建了一个绑定路由的web引擎
func NewHttpEngine() (*gin.Engine, error) {
	//设置为Release，为的是默认在启动中不输出调试信息
	gin.SetMode(gin.ReleaseMode)
	//默认启动一个web引擎

	r := gin.Default()

	//业务绑定路由操作
	Routes(r)

	return r, nil
}
