package contract

import "net/http"

//提供kernel服务凭证
const KernelKey = "hade:kernel"

//Kernel 接口提供最核心的结构
type Kernel interface {
	//HttpEngine http.Handler结构，作为net/http框架使用, 实际上是gin.Engine
	HttpEngine() http.Handler
}
