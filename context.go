package web

import (
	"codnect.io/procyon-web/http"
	"context"
	stdhttp "net/http"
	"time"
)

type defaultHttpServerContext struct {
	parent   *defaultHttpServerContext
	context  context.Context
	request  defaultHttpServerRequest
	response defaultHttpServerResponse

	HandlerChain     http.HandlerChain
	nextHandlerIndex int

	err       error
	completed bool
	aborted   bool

	delegate      defaultRequestDelegate
	pathVariables http.PathVariables
}

func (c *defaultHttpServerContext) WithValue(key, val any) http.Context {
	copyContext := new(defaultHttpServerContext)
	*copyContext = *c

	ctx := c.context
	if ctx == nil {
		ctx = context.Background()
	}

	copyContext.context = context.WithValue(ctx, key, val)
	return copyContext
}

func (c *defaultHttpServerContext) With(request http.Request, response http.Response) http.Context {
	if request == nil {
		panic("nil request")
	}

	if response == nil {
		panic("nil response")
	}

	copyContext := new(defaultHttpServerContext)
	*copyContext = *c
	copyContext.request = *(request.(*defaultHttpServerRequest))
	copyContext.response = *(response.(*defaultHttpServerResponse))

	if c.parent == nil {
		copyContext.parent = c
	}

	return copyContext
}

func (c *defaultHttpServerContext) WithRequest(request http.Request) http.Context {
	if request == nil {
		panic("nil request")
	}

	copyContext := new(defaultHttpServerContext)
	*copyContext = *c
	copyContext.request = *(request.(*defaultHttpServerRequest))

	if c.parent == nil {
		copyContext.parent = c
	}

	return copyContext
}

func (c *defaultHttpServerContext) WithResponse(response http.Response) http.Context {
	if response == nil {
		panic("nil response")
	}

	copyContext := new(defaultHttpServerContext)
	*copyContext = *c
	copyContext.response = *(response.(*defaultHttpServerResponse))

	if c.parent == nil {
		copyContext.parent = c
	}

	return copyContext
}

func (c *defaultHttpServerContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *defaultHttpServerContext) Done() <-chan struct{} {
	return nil
}

func (c *defaultHttpServerContext) setErr(err error) {
	if c.parent != nil {
		c.parent.setErr(err)
	} else {
		c.err = err
	}
}

func (c *defaultHttpServerContext) Err() error {
	if c.parent != nil {
		return c.parent.Err()
	}

	return c.err
}

func (c *defaultHttpServerContext) Value(key any) any {
	if key == http.PathVariablesAttribute {
		return &c.pathVariables
	}

	if c.context == nil {
		return nil
	}

	return c.context.Value(key)
}

func (c *defaultHttpServerContext) Parent() http.Context {
	return c.parent
}

func (c *defaultHttpServerContext) complete() {
	if c.parent != nil {
		c.parent.complete()
	} else {
		c.completed = true
	}
}

func (c *defaultHttpServerContext) IsCompleted() bool {
	if c.parent != nil {
		return c.parent.IsCompleted()
	}

	return c.completed
}

func (c *defaultHttpServerContext) Abort() {
	if c.parent != nil {
		c.parent.Abort()
	} else {
		c.aborted = true
	}
}

func (c *defaultHttpServerContext) IsAborted() bool {
	if c.parent != nil {
		return c.parent.IsAborted()
	}

	return c.aborted
}

func (c *defaultHttpServerContext) Request() http.Request {
	return &c.request
}

func (c *defaultHttpServerContext) Response() http.Response {
	return &c.response
}

func (c *defaultHttpServerContext) Reset(req *stdhttp.Request, writer stdhttp.ResponseWriter) {
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

func (c *defaultHttpServerContext) nextHandler() int {
	if c.parent != nil {
		return c.parent.nextHandler()
	}

	return c.nextHandlerIndex
}

func (c *defaultHttpServerContext) setNextHandler(nextHandler int) {
	if c.parent != nil {
		c.parent.setNextHandler(nextHandler)
	} else {
		c.nextHandlerIndex = nextHandler
	}
}

func (c *defaultHttpServerContext) Invoke(ctx http.Context) {
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
