package middleware

import (
	"github.com/jamesluo111/core_web/framework/gin"
	"log"
	"time"
)

//统计请求时长中间件
func StatDuration() gin.HandlerFunc {
	return func(c *framework.Context) error {
		nowPre := time.Now()
		c.Next()
		nowSuf := time.Now()
		duration := nowSuf.Sub(nowPre)
		log.Println(c.GetRequest().URL, c.GetRequest().Method, duration.String())
		return nil
	}
}
