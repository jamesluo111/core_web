package demo

import (
	"github.com/jamesluo111/core_web/framework"
)

// 服务提供方
type DemoProvider struct {
	framework.ServiceProvider

	c framework.Container
}

// Name方法直接将服务对应的字符串凭证返回, 这个例子就是hade.demo
func (sp *DemoProvider) Name() string {
	return DemoKey
}

// Register 方法是注册初始化服务实例的方法，这里先暂定为 NewDemoService
func (sp *DemoProvider) Register(c framework.Container) framework.NewInstance {
	return NewDemoService
}

// IsDefer 方法表示是否延迟实例化，我们这里设置为 true，将这个服务的实例化延迟到第一次 make 的时候
func (sp *DemoProvider) IsDefer() bool {
	return false
}

// Params 方法表示实例化的参数。我们这里只实例化一个参数：container，表示我们在 NewDemoService 这个函数中，只有一个参数，container
func (sp *DemoProvider) Params(c framework.Container) []interface{} {
	return []interface{}{sp.c}
}

// Boot 方法我们这里我们什么逻辑都不执行, 只打印一行日志信息
func (sp *DemoProvider) Boot(c framework.Container) error {
	sp.c = c
	return nil
}
