package demo

import (
	"fmt"
	"github.com/jamesluo111/core_web/framework"
)

// 具体的接口实例
type DemoService struct {
	// 实现接口
	Service

	// 参数
	c framework.Container
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	//这里需要将参数展开
	c := params[0].(framework.Container)

	fmt.Println("new demo service")

	return &DemoService{c: c}, nil
}

// 实现接口
func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "i am foo",
	}
}
