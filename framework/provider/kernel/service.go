package kernel

import (
	"github.com/jamesluo111/core_web/framework/gin"
	"net/http"
)

type HadeKernelService struct {
	engine *gin.Engine
}

//初始化web服务引擎实例
func NewHadeKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &HadeKernelService{engine: httpEngine}, nil
}

func (s *HadeKernelService) HttpEngine() http.Handler {
	return s.engine
}
