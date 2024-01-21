package server

import (
	"codnect.io/logy"
	"codnect.io/procyon-web/http"
	"codnect.io/procyon-web/http/router"
	"context"
	"errors"
	"fmt"
	stdhttp "net/http"
	"sync"
)

var (
	log = logy.Get()
)

type Server struct {
	props  http.ServerProperties
	server *stdhttp.Server

	contextPool sync.Pool

	mappingRegistry *router.MappingRegistry
	errorHandler    http.ErrorHandler
}

func New() *Server {
	return &Server{
		contextPool: sync.Pool{
			New: func() any {
				return newServerContext()
			},
		},
		props: http.ServerProperties{
			Port: 8080,
		},
	}
}

func (s *Server) Start() error {
	s.server = &stdhttp.Server{
		Addr:    fmt.Sprintf(":%d", s.props.Port),
		Handler: s,
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			log.Info("Server started on port(s): {} (http)", s.props.Port)
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Server) Port() int {
	return s.props.Port
}

func (s *Server) ServeHTTP(writer stdhttp.ResponseWriter, request *stdhttp.Request) {
	ctx := s.contextPool.Get().(*Context)
	ctx.Reset(request, writer)

	defer func() {
		if runtimeError := recover(); runtimeError != nil {
			s.handlerRuntimeError(ctx, runtimeError)
			s.contextPool.Put(ctx)
		}
	}()

	chain, exists := s.mappingRegistry.GetHandler(ctx)

	if exists {
		ctx.HandlerChain = chain
		ctx.Invoke(ctx)

		if ctx.Err() != nil {
			s.errorHandler.HandleError(ctx, ctx.Err())
		}
	} else {
		s.errorHandler.HandleError(ctx, &http.NotFoundError{})
	}

	s.contextPool.Put(ctx)
}

func (s *Server) handlerRuntimeError(ctx http.Context, runtimeError any) {
	switch err := runtimeError.(type) {
	case error:
		s.errorHandler.HandleError(ctx, err)
	case string:
		s.errorHandler.HandleError(ctx, errors.New(err))
	default:
		s.errorHandler.HandleError(ctx, fmt.Errorf("unknown error: %v", err))
	}
}
