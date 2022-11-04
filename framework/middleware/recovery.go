package middleware

import "jamesluo1/framework"

//recovery机制，将协程中的异常捕获
func Recovery() framework.ControllerHandler {
	//使用函数回调
	return func(c *framework.Context) error {
		//核心在增加这个recovery机制,捕获c.Next()出现的panic
		defer func() {
			if err := recover(); err != nil {
				c.Json(500, err)
				return
			}
		}()
		c.Next()

		return nil
	}
}
