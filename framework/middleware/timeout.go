package middleware

import (
	"context"
	"fmt"
	"jamesluo1/framework"
	"log"
	"time"
)

func TimeoutHandler(fun framework.ControllerHandler, d time.Duration) framework.ControllerHandler {
	//使用函数回调
	return func(c *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		//执行业务逻辑前预操作：初始化超时context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		//c.request.WithContext(durationCtx)

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			//执行具体业务逻辑
			fun(c)
			finish <- struct{}{}
		}()

		//执行业务逻辑后的操作
		select {
		case p := <-panicChan:
			log.Println(p)
			//c.response.WriteHeader(500)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			//c.response.Write([]byte("time out"))
		}
		return nil
	}
}