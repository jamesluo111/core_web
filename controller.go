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
			c.SetOkStatus().Json("ok")

			finish <- struct{}{}
		}()
	}()

	select {
	case p := <-pinacChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.SetStatus(500).Json("panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.SetStatus(500).Json("time out")
		c.SetHasTimeout()
	}
	return nil
}
