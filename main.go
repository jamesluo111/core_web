package main

import (
	"encoding/json"
	"jamesluo1/framework"
	"jamesluo1/framework/middleware"
	"net/http"
	"strconv"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Test1(), middleware.Test2())
	registerRouter(core)
	server := &http.Server{
		Addr:    ":8888",
		Handler: core,
	}
	server.ListenAndServe()
}

type FooServer struct {
}

func (f FooServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	obj := map[string]interface{}{
		"data": nil,
	}
	//设置控制器response的header
	res.Header().Set("Content-Type", "application/json")
	//从请求体中获取参数
	foo := req.PostFormValue("foo")
	if foo == "" {
		foo = "10"
	}
	fooInt, err := strconv.Atoi(foo)
	if err != nil {
		res.WriteHeader(500)
		return
	}
	obj["data"] = fooInt
	byt, err := json.Marshal(obj)
	if err != nil {
		res.WriteHeader(500)
		return
	}
	res.WriteHeader(200)
	res.Write(byt)
	return
}
