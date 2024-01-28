package http

import (
	"codnect.io/procyon-core/runtime/env/property"
	"context"
)

type Server interface {
	Start() error
	Stop(ctx context.Context) error
	Port() int
}

type ServerProperties struct {
	property.Properties `prefix:"server"`

	Port int `prop:"port" default:"8080"`
}
