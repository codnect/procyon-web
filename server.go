package web

import (
	"codnect.io/procyon-core/runtime"
	"codnect.io/procyon-web/http"
	"codnect.io/procyon-web/http/router"
	"context"
	"errors"
	"fmt"
	stdhttp "net/http"
	"sync"
)

type defaultHttpServer struct {
	server      *stdhttp.Server
	contextPool sync.Pool

	mappingRegistry *router.MappingRegistry
	errorHandler    http.ErrorHandler

	props   runtime.ServerProperties
	handler http.Handler
}

func NewDefaultHttpServer(handler http.Handler) *defaultHttpServer {
	return &defaultHttpServer{
		contextPool: sync.Pool{
			New: func() any {
				return &defaultHttpServerContext{
					pathVariables: http.PathVariables{},
					response: defaultHttpServerResponse{
						headers: map[string][]string{},
					},
				}
			},
		},
		props: runtime.ServerProperties{
			Port: 8080,
		},
		handler: handler,
	}
}

func (s *defaultHttpServer) Start(ctx context.Context) error {
	s.server = &stdhttp.Server{
		Addr:    fmt.Sprintf(":%d", s.props.Port),
		Handler: s,
	}

	if err := s.server.ListenAndServe(); err != nil {
		//log.Info("DefaultServer started on port(s): {} (http)", s.props.Port)
	}

	return nil
}

func (s *defaultHttpServer) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (s *defaultHttpServer) Port() int {
	return s.props.Port
}

func (s *defaultHttpServer) ServeHTTP(writer stdhttp.ResponseWriter, request *stdhttp.Request) {
	ctx := s.contextPool.Get().(*defaultHttpServerContext)
	ctx.Reset(request, writer)

	s.handler.Invoke(ctx)
	/*
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
		}*/

	s.contextPool.Put(ctx)
}

func (s *defaultHttpServer) handlerRuntimeError(ctx http.Context, runtimeError any) {
	switch err := runtimeError.(type) {
	case error:
		s.errorHandler.HandleError(ctx, err)
	case string:
		s.errorHandler.HandleError(ctx, errors.New(err))
	default:
		s.errorHandler.HandleError(ctx, fmt.Errorf("unknown error: %v", err))
	}
}
