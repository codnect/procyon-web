package http

import (
	"context"
	"time"
)

type Context interface {
	context.Context

	IsCompleted() bool
	Abort()
	IsAborted() bool
	Request() Request
	Response() Response
}

type contextWrapper struct {
	parent          Context
	requestWrapper  requestWrapper
	responseWrapper responseWrapper
	key             any
	value           any
}

func (c *contextWrapper) Deadline() (deadline time.Time, ok bool) {
	return c.parent.Deadline()
}

func (c *contextWrapper) Done() <-chan struct{} {
	return c.parent.Done()
}

func (c *contextWrapper) Err() error {
	return c.parent.Err()
}

func (c *contextWrapper) Value(key any) any {
	if c.key != nil && c.key == key {
		return c.value
	}

	return c.parent.Value(key)
}

func (c *contextWrapper) IsCompleted() bool {
	return c.parent.IsCompleted()
}

func (c *contextWrapper) Abort() {
	c.parent.Abort()
}

func (c *contextWrapper) IsAborted() bool {
	return c.parent.IsAborted()
}

func (c *contextWrapper) Request() Request {
	return c.requestWrapper
}

func (c *contextWrapper) Response() Response {
	return c.responseWrapper
}

func NewContext(request Request, response Response) Context {
	if request == nil {
		panic("nil request")
	}

	if response == nil {
		panic("nil response")
	}

	context := &contextWrapper{
		requestWrapper: requestWrapper{
			request: request,
		},
		responseWrapper: responseWrapper{
			response: response,
		},
	}

	context.requestWrapper.context = context
	context.responseWrapper.context = context
	return context
}

func WithValue(parent Context, key, val any) Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}

	if key == nil {
		panic("nil key")
	}

	context := &contextWrapper{
		parent: parent,
		requestWrapper: requestWrapper{
			request: parent.Request(),
		},
		responseWrapper: responseWrapper{
			response: parent.Response(),
		},
		key:   key,
		value: val,
	}

	context.requestWrapper.context = context
	context.responseWrapper.context = context
	return context
}

func WithRequest(parent Context, request Request) Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}

	if request == nil {
		panic("nil request")
	}

	context := &contextWrapper{
		parent: parent,
		requestWrapper: requestWrapper{
			request: request,
		},
		responseWrapper: responseWrapper{
			response: parent.Response(),
		},
	}

	context.requestWrapper.context = context
	context.responseWrapper.context = context
	return context
}

func WithResponse(parent Context, response Response) Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}

	if response == nil {
		panic("nil response")
	}

	context := &contextWrapper{
		parent: parent,
		requestWrapper: requestWrapper{
			request: parent.Request(),
		},
		responseWrapper: responseWrapper{
			response: response,
		},
	}

	context.requestWrapper.context = context
	context.responseWrapper.context = context
	return context
}
