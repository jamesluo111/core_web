package middleware

import (
	"jamesluo1/framework"
	"log"
	"time"
)

//统计请求时长中间件
func StatDuration() framework.ControllerHandler {
	return func(c *framework.Context) error {
		nowPre := time.Now()
		c.Next()
		nowSuf := time.Now()
		duration := nowSuf.Sub(nowPre)
		log.Println(c.GetRequest().URL, c.GetRequest().Method, duration.String())
		return nil
	}
}
