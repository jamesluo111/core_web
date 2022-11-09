package demo

import (
	"fmt"
	"github.com/jamesluo111/core_web/framework"
)

// 具体的接口实例
type Service struct {
	// 参数
	c framework.Container
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	//这里需要将参数展开
	c := params[0].(framework.Container)

	fmt.Println("new demo service")

	return &Service{c: c}, nil
}

// 实现接口
func (s *Service) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "foo",
		},
		{
			ID:   2,
			Name: "bar",
		},
	}
}
