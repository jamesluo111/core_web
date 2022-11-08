package middleware

import (
	"context"
	"fmt"
	"github.com/jamesluo111/core_web/framework/gin"
	"log"
	"time"
)

func TimeoutHandler(d time.Duration) gin.HandlerFunc {
	//使用函数回调
	return func(c *gin.Context) {
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
			c.Next()
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
			c.ISetStatus(500).IJson("time out")
			//c.response.Write([]byte("time out"))
		}
		return
	}
}
