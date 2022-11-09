package main

import (
	"github.com/jamesluo111/core_web/app/console"
	"github.com/jamesluo111/core_web/app/http"
	"github.com/jamesluo111/core_web/framework"
	"github.com/jamesluo111/core_web/framework/provider/app"
	"github.com/jamesluo111/core_web/framework/provider/kernel"
)

func main() {
	container := framework.NewHadeContainer()
	//绑定APP服务提供者
	container.Bind(&app.HadeAppProvider{})
	//将http引擎初始化，并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}
	console.RunCommand(container)
	//core := gin.New()
	//core.Bind(&demo.DemoServiceProvider{})
	//core.Use(gin.Recovery())
	//core.Use(middleware.Cost())
	//registerRouter(core)
	//server := &http.Server{
	//	Addr:    ":8888",
	//	Handler: core,
	//}
	//go func() {
	//	server.ListenAndServe()
	//}()
	//当前的Goruoutine等待信号量
	//quit := make(chan os.Signal)
	////监控信号：SIGINT, SIGTERM, SIGQUIT
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	////这里会阻塞当前Goruoutine等待信号
	//<-quit
	//
	//// 调用Server.Shutdown graceful结束
	//cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//if err := server.Shutdown(cxt); err != nil {
	//	log.Fatal("server shutdown：", err)
	//}
	//defer cancel()
}
