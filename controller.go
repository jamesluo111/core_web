package main

import (
	"context"
	"fmt"
	"jamesluo1/framework"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	pinacChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(12*time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				pinacChan <- p
			}
			time.Sleep(10 * time.Second)
			c.Json(200, "ok")

			finish <- struct{}{}
		}()
	}()

	select {
	case p := <-pinacChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}
