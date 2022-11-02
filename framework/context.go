package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Context struct {
	request  *http.Request
	response http.ResponseWriter
	ctx      context.Context

	//标记是否超时
	hasTimeOut bool
	//写保护机制
	writerMux *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:   r,
		response:  w,
		ctx:       r.Context(),
		writerMux: &sync.Mutex{},
	}
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.response
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeOut = true
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		lens := len(vals)
		if lens > 0 {
			intval, err := strconv.Atoi(vals[lens-1])
			if err != nil {
				return def
			}
			return intval
		}
	}
	return def
}

func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		lens := len(vals)
		if lens > 0 {
			return vals[lens-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

func (ctx *Context) FromInt(key string, def int) int {
	params := ctx.FromAll()
	if vals, ok := params[key]; ok {
		lens := len(vals)
		if lens > 0 {
			intval, err := strconv.Atoi(vals[lens-1])
			if err != nil {
				return def
			}
			return intval
		}
	}
	return def
}

func (ctx *Context) FromString(key string, def string) string {
	params := ctx.FromAll()
	if vals, ok := params[key]; ok {
		lens := len(vals)
		if lens > 0 {
			return vals[lens-1]
		}
	}
	return def
}

func (ctx *Context) FromArray(key string, def []string) []string {
	params := ctx.FromAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) FromAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.hasTimeOut {
		return nil
	}
	ctx.response.Header().Set("Content-Type", "application/json")
	ctx.response.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.response.WriteHeader(500)
		return err
	}
	ctx.response.Write(byt)
	return nil
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}