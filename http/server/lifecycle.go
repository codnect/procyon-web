package server

type StartStopLifecycle struct {
	server  Server
	running bool
}

func (l *StartStopLifecycle) Start() error {
	err := l.server.Start()

	if err != nil {
		return err
	}

	l.running = true
	return nil
}

func (l *StartStopLifecycle) Stop() error {
	err := l.server.Stop()

	if err != nil {
		return err
	}

	l.running = false
	return nil
}

func (l *StartStopLifecycle) IsRunning() bool {
	return l.running
}
