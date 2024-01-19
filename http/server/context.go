package server

import (
	"codnect.io/procyon-web/http"
	"context"
	stdhttp "net/http"
	"time"
)

type Context struct {
	parent   *Context
	context  context.Context
	request  Request
	response Response

	HandlerChain     http.HandlerChain
	nextHandlerIndex int

	err       error
	completed bool
	aborted   bool

	delegate      ContextDelegate
	pathVariables http.PathVariables
}

func newServerContext() *Context {
	return &Context{
		pathVariables: http.PathVariables{},
	}
}

func (c *Context) WithValue(key, val any) http.Context {
	copyContext := new(Context)
	*copyContext = *c

	ctx := c.context
	if ctx == nil {
		ctx = context.Background()
	}

	copyContext.context = context.WithValue(ctx, key, val)
	return copyContext
}

func (c *Context) With(request http.Request, response http.Response) http.Context {
	if request == nil {
		panic("nil request")
	}

	if response == nil {
		panic("nil response")
	}

	copyContext := new(Context)
	*copyContext = *c
	copyContext.request = *(request.(*Request))
	copyContext.response = *(response.(*Response))

	if c.parent == nil {
		copyContext.parent = c
	}

	return copyContext
}

func (c *Context) WithRequest(request http.Request) http.Context {
	if request == nil {
		panic("nil request")
	}

	copyContext := new(Context)
	*copyContext = *c
	copyContext.request = *(request.(*Request))

	if c.parent == nil {
		copyContext.parent = c
	}

	return copyContext
}

func (c *Context) WithResponse(response http.Response) http.Context {
	if response == nil {
		panic("nil response")
	}

	copyContext := new(Context)
	*copyContext = *c
	copyContext.response = *(response.(*Response))

	if c.parent == nil {
		copyContext.parent = c
	}

	return copyContext
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *Context) Done() <-chan struct{} {
	return nil
}

func (c *Context) setErr(err error) {
	if c.parent != nil {
		c.parent.setErr(err)
	} else {
		c.err = err
	}
}

func (c *Context) Err() error {
	if c.parent != nil {
		return c.parent.Err()
	}

	return c.err
}

func (c *Context) Value(key any) any {
	if key == http.PathVariablesAttribute {
		return &c.pathVariables
	}

	if c.context == nil {
		return nil
	}

	return c.context.Value(key)
}

func (c *Context) Parent() *Context {
	return c.parent
}

func (c *Context) complete() {
	if c.parent != nil {
		c.parent.complete()
	} else {
		c.completed = true
	}
}

func (c *Context) IsCompleted() bool {
	if c.parent != nil {
		return c.parent.IsCompleted()
	}

	return c.completed
}

func (c *Context) Abort() {
	if c.parent != nil {
		c.parent.Abort()
	} else {
		c.aborted = true
	}
}

func (c *Context) IsAborted() bool {
	if c.parent != nil {
		return c.parent.IsAborted()
	}

	return c.aborted
}

func (c *Context) Request() http.Request {
	return &c.request
}

func (c *Context) Response() http.Response {
	return &c.response
}

func (c *Context) Reset(req *stdhttp.Request, writer stdhttp.ResponseWriter) {
	/*if !c.IsCompleted() {
		return
	}*/

	c.request.req = req
	c.response.writer = writer
	c.delegate.ctx = c

	c.parent = nil
	c.err = nil
	c.context = nil
	c.completed = false
	c.aborted = false

	c.nextHandlerIndex = 0
	//c.pathVariables.currentIndex = 0
}

func (c *Context) nextHandler() int {
	if c.parent != nil {
		return c.parent.nextHandler()
	}

	return c.nextHandlerIndex
}

func (c *Context) setNextHandler(nextHandler int) {
	if c.parent != nil {
		c.parent.setNextHandler(nextHandler)
	} else {
		c.nextHandlerIndex = nextHandler
	}
}

func (c *Context) Invoke(ctx http.Context) {
	if len(c.HandlerChain) == 0 {
		return
	}

	nextHandler := c.nextHandler()

	if c.IsCompleted() || c.IsAborted() || len(c.HandlerChain) <= nextHandler {
		return
	}

	next := c.HandlerChain[nextHandler]
	nextHandler++
	c.setNextHandler(nextHandler)

	err := next(ctx, c.delegate)

	if err != nil {
		c.setErr(err)
	}

	if c.IsCompleted() || c.IsAborted() {
		return
	}

	if nextHandler != len(c.HandlerChain) {
		c.Abort()
	} else {
		c.complete()
	}
}
