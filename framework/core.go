package framework

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]map[string]ControllerHandler
}

func NewCore() *Core {
	//定义二级map
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}
	//将二级map写入到一级map中
	router := map[string]map[string]ControllerHandler{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter
	return &Core{router: router}
}

//对应get
func (c Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler
}

func (c Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
}

func (c Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
}

func (c Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = handler
}

func (c Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	upperUri := strings.ToUpper(uri)
	if methodHandler, ok := c.router[upperMethod]; ok {
		if handler, ok := methodHandler[upperUri]; ok {
			return handler
		}
	}
	return nil
}

func (c Core) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("core.ServeHTTP")
	ctx := NewContext(req, res)
	// 一个简单的路由选择器，这里直接写死为测试路由foo
	fmt.Println(req.URL.Path)
	router := c.FindRouteByRequest(req)
	if router == nil {
		ctx.Json(404, "not found")
		return
	}
	log.Println("core.router")
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
