package framework

import (
	"fmt"
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c Core) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("core.ServeHTTP")
	ctx := NewContext(req, res)
	// 一个简单的路由选择器，这里直接写死为测试路由foo
	fmt.Println(req.URL.Path)
	router := c.router[req.URL.Path]
	if router == nil {
		res.WriteHeader(404)
		return
	}
	log.Println("core.router")
	router(ctx)
}
