package http

import (
	"time"
)

type Cookie struct {
	Name  string
	Value string

	Path    string
	Domain  string
	Expires time.Time

	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite SameSite
}

type SameSite int

const (
	SameSiteDefaultMode SameSite = iota + 1
	SameSiteLaxMode
	SameSiteStrictMode
	SameSiteNoneMode
)
