package web

import (
	"codnect.io/procyon-web/http"
	"context"
)

type serverLifecycle struct {
	server  http.Server
	running bool
}

func newServerLifecycle(server http.Server) *serverLifecycle {
	return &serverLifecycle{}
}

func (l *serverLifecycle) Start(ctx context.Context) error {
	err := l.server.Start()

	if err != nil {
		return err
	}

	l.running = true
	return nil
}

func (l *serverLifecycle) Stop(ctx context.Context) error {
	err := l.server.Stop(ctx)

	if err != nil {
		return err
	}

	l.running = false
	return nil
}

func (l *serverLifecycle) IsRunning() bool {
	return l.running
}
