package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler // 从core这边设置的中间件
}

func NewCore() *Core {
	//将二级map写入到一级map中
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

//对应get
func (c *Core) Get(url string, handlers ...ControllerHandler) {
	allhandler := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allhandler...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allhandler := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allhandler...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allhandler := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allhandler...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allhandler := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allhandler...); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouteNodeByRequest(request *http.Request) *node {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	if methodHandler, ok := c.router[upperMethod]; ok {
		return methodHandler.root.matchNode(uri)
	}
	return nil
}

func (c *Core) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("core.ServeHTTP")
	ctx := NewContext(req, res)
	//寻找路由
	node := c.FindRouteNodeByRequest(req)
	if node == nil {
		ctx.Json(404, "not found")
		return
	}
	//设置context中的handler字段
	ctx.SetHandlers(node.handlers)

	//设置路由参数
	params := node.parseParamsFromEndNode(req.URL.Path)
	ctx.SetParams(params)
	log.Println("core.router")
	if err := ctx.Next(); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}

func (c *Core) Use(middleware ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middleware...)
}
