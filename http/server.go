package http

import "codnect.io/procyon-core/env/property"

type Server interface {
	Start() error
	Stop() error
	Port() int
	ShutDownGracefully(ctx Context) error
}

type ServerProperties struct {
	property.Properties `prefix:"server"`

	Port int `prop:"port" default:"8080"`
}
