package main

import (
	"context"
	"github.com/jamesluo111/core_web/framework/gin"
	"github.com/jamesluo111/core_web/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()
	core.Use(gin.Recovery())
	core.Use(middleware.Cost())
	registerRouter(core)
	server := &http.Server{
		Addr:    ":8888",
		Handler: core,
	}
	go func() {
		server.ListenAndServe()
	}()
	//当前的Goruoutine等待信号量
	quit := make(chan os.Signal)
	//监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//这里会阻塞当前Goruoutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := server.Shutdown(cxt); err != nil {
		log.Fatal("server shutdown：", err)
	}
	defer cancel()
}
