package middleware

import (
	"fmt"
	"github.com/jamesluo111/core_web/framework/gin"
)

func Test1() gin.HandlerFunc {
	//使用函数回调
	return func(c *gin.Context) {
		fmt.Println("middleware pre test1")
		c.Next() //调用Next往下调用，会自增context.index
		fmt.Println("middleware post test1")
		return
	}
}

func Test2() gin.HandlerFunc {
	//使用函数回调
	return func(c *gin.Context) {
		fmt.Println("middleware pre test2")
		c.Next() //调用Next往下调用，会自增context.index
		fmt.Println("middleware post test2")
		return
	}
}

func Test3() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
		return
	}
}
