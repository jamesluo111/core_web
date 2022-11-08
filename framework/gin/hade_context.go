// Copyright 2021 jianfengye.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"context"
	"github.com/jamesluo111/core_web/framework"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

// engine实现container的封装绑定
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

// IsBind关键字凭证是否已经绑定服务提供者
func (engin *Engine) IsBind(key string) bool {
	return engin.container.IsBind(key)
}

//context实现container的封装
func (ctx *Context) Make(key string) (interface{}, error) {
	return ctx.container.Make(key)
}

func (ctx *Context) MustMake(key string) interface{} {
	return ctx.container.MustMake(key)
}

func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ctx.container.MakeNew(key, params)
}
