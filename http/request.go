package http

import (
	"io"
)

type Request interface {
	Context() Context

	Cookie(name string) (*Cookie, bool)
	Cookies() []*Cookie

	Parameter(name string) (string, bool)
	ParameterNames() []string
	Parameters(name string) []string
	QueryString() string

	Header(name string) (string, bool)
	HeaderNames() []string
	Headers(name string) []string

	Path() string
	Method() Method
	Reader() io.Reader
	Scheme() string
	IsSecure() bool
}

type requestWrapper struct {
	request Request
	context Context
}

func (r requestWrapper) Context() Context {
	return r.context
}

func (r requestWrapper) Cookie(name string) (*Cookie, bool) {
	return r.request.Cookie(name)
}

func (r requestWrapper) Cookies() []*Cookie {
	return r.request.Cookies()
}

func (r requestWrapper) Parameter(name string) (string, bool) {
	return r.request.Parameter(name)
}

func (r requestWrapper) ParameterNames() []string {
	return r.request.ParameterNames()
}

func (r requestWrapper) Parameters(name string) []string {
	return r.request.Parameters(name)
}

func (r requestWrapper) QueryString() string {
	return r.request.QueryString()
}

func (r requestWrapper) Header(name string) (string, bool) {
	return r.request.Header(name)
}

func (r requestWrapper) HeaderNames() []string {
	return r.request.HeaderNames()
}

func (r requestWrapper) Headers(name string) []string {
	return r.request.Headers(name)
}

func (r requestWrapper) Path() string {
	return r.request.Path()
}

func (r requestWrapper) Method() Method {
	return r.request.Method()
}

func (r requestWrapper) Reader() io.Reader {
	return r.request.Reader()
}

func (r requestWrapper) Scheme() string {
	return r.request.Scheme()
}

func (r requestWrapper) IsSecure() bool {
	return r.request.IsSecure()
}

type MultiReadRequest struct {
	request Request
	reader  io.Reader
}

func NewMultiReadRequest(request Request) *MultiReadRequest {
	if request == nil {
		panic("request cannot be nil")
	}

	return &MultiReadRequest{
		request: request,
	}
}

func (m *MultiReadRequest) Context() Context {
	return m.request.Context()
}

func (m *MultiReadRequest) Cookie(name string) (*Cookie, bool) {
	return m.request.Cookie(name)
}

func (m *MultiReadRequest) Cookies() []*Cookie {
	return m.request.Cookies()
}

func (m *MultiReadRequest) Parameter(name string) (string, bool) {
	return m.request.Parameter(name)
}

func (m *MultiReadRequest) ParameterNames() []string {
	return m.request.ParameterNames()
}

func (m *MultiReadRequest) Parameters(name string) []string {
	return m.request.Parameters(name)
}

func (m *MultiReadRequest) QueryString() string {
	return m.request.QueryString()
}

func (m *MultiReadRequest) Header(name string) (string, bool) {
	return m.request.Header(name)
}

func (m *MultiReadRequest) HeaderNames() []string {
	return m.request.HeaderNames()
}

func (m *MultiReadRequest) Headers(name string) []string {
	return m.request.Headers(name)
}

func (m *MultiReadRequest) Path() string {
	return m.request.Path()
}

func (m *MultiReadRequest) Method() Method {
	return m.request.Method()
}

func (m *MultiReadRequest) Reader() io.Reader {
	if m.reader == nil {
		m.reader = newCachedReader(m.request.Reader())
	}

	return m.reader
}

func (m *MultiReadRequest) Scheme() string {
	return m.request.Scheme()
}

func (m *MultiReadRequest) IsSecure() bool {
	return m.request.IsSecure()
}

type cachedReader struct {
	data   []byte
	reader io.Reader
}

func newCachedReader(reader io.Reader) *cachedReader {
	return &cachedReader{
		data:   nil,
		reader: reader,
	}
}

func (c *cachedReader) readAll() (err error) {
	if c.data == nil {
		c.data, err = io.ReadAll(c.reader)
	}

	return err
}

func (c *cachedReader) Read(p []byte) (n int, err error) {
	err = c.readAll()
	if err != nil {
		return 0, err
	}

	n = copy(p, c.data)
	return
}
