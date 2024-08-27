package web

import (
	"codnect.io/procyon-core/runtime"
	"context"
)

type serverLifecycle struct {
	server  runtime.Server
	running bool
}

func newServerLifecycle(server runtime.Server) *serverLifecycle {
	return &serverLifecycle{
		server: server,
	}
}

func (l *serverLifecycle) Start(ctx context.Context) error {
	err := l.server.Start(ctx)

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
