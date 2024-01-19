package server

import (
	"codnect.io/procyon-web/http"
	"io"
	stdhttp "net/http"
)

type Request struct {
	req    *stdhttp.Request
	ctx    *Context
	reader io.Reader
}

func (r *Request) WithReader(reader io.Reader) http.Request {
	if reader == nil {
		panic("nil reader")
	}

	copyRequest := new(Request)
	*copyRequest = *r
	copyRequest.reader = reader
	return copyRequest
}

func (r *Request) Context() http.Context {
	return r.ctx
}

func (r *Request) Cookie(name string) (*http.Cookie, bool) {
	return nil, false
}

func (r *Request) Cookies() []*http.Cookie {
	return nil
}

func (r *Request) QueryParameter(name string) (string, bool) {
	return "", false
}

func (r *Request) QueryParameterNames() []string {
	return nil
}

func (r *Request) QueryParameters(name string) []string {
	return nil
}

func (r *Request) QueryString() string {
	return ""
}

func (r *Request) Header(name string) (string, bool) {
	return "", false
}

func (r *Request) HeaderNames() []string {
	return nil
}

func (r *Request) Headers(name string) []string {
	return nil
}

func (r *Request) Path() string {
	return r.req.URL.Path
}

func (r *Request) Method() http.Method {
	return http.Method(r.req.Method)
}

func (r *Request) Reader() io.Reader {
	if r.reader != nil {
		return r.reader
	}

	return r.req.Body
}

func (r *Request) Scheme() string {
	return ""
}

func (r *Request) IsSecure() bool {
	return false
}
