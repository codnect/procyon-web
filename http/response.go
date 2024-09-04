package http

import (
	"io"
)

type Response interface {
	Context() Context
	AddCookie(cookie *Cookie)

	ContentLength() int
	SetContentLength(len int)

	CharacterEncoding() string
	SetCharacterEncoding(charset string)

	ContentType() string
	SetContentType(contentType string)

	AddHeader(name string, value string)
	SetHeader(name string, value string)
	DeleteHeader(name string)
	Header(name string) (string, bool)
	HeaderNames() []string
	Headers(name string) []string
	Status() Status
	SetStatus(status Status)

	Writer() io.Writer
	Flush() error
	IsCommitted() bool
	Reset()
}

type responseWrapper struct {
	response Response
	context  Context
}

func (r responseWrapper) Context() Context {
	return r.context
}

func (r responseWrapper) AddCookie(cookie *Cookie) {
	r.response.AddCookie(cookie)
}

func (r responseWrapper) ContentLength() int {
	return r.response.ContentLength()
}

func (r responseWrapper) SetContentLength(len int) {
	r.response.SetContentLength(len)
}

func (r responseWrapper) CharacterEncoding() string {
	return r.response.CharacterEncoding()
}

func (r responseWrapper) SetCharacterEncoding(charset string) {
	r.response.SetCharacterEncoding(charset)
}

func (r responseWrapper) ContentType() string {
	return r.response.ContentType()
}

func (r responseWrapper) SetContentType(contentType string) {
	r.response.SetContentType(contentType)
}

func (r responseWrapper) AddHeader(name string, value string) {
	r.response.AddHeader(name, value)
}

func (r responseWrapper) SetHeader(name string, value string) {
	r.response.SetHeader(name, value)
}

func (r responseWrapper) DeleteHeader(name string) {
	r.response.DeleteHeader(name)
}

func (r responseWrapper) Header(name string) (string, bool) {
	return r.response.Header(name)
}

func (r responseWrapper) HeaderNames() []string {
	return r.response.HeaderNames()
}

func (r responseWrapper) Headers(name string) []string {
	return r.response.Headers(name)
}

func (r responseWrapper) Status() Status {
	return r.response.Status()
}

func (r responseWrapper) SetStatus(status Status) {
	r.response.SetStatus(status)
}

func (r responseWrapper) Writer() io.Writer {
	return r.response.Writer()
}

func (r responseWrapper) Flush() error {
	return r.response.Flush()
}

func (r responseWrapper) IsCommitted() bool {
	return r.response.IsCommitted()
}

func (r responseWrapper) Reset() {
	r.response.Reset()
}

type MultiReadResponse struct {
	response Response
}

func NewMultiReadResponse(response Response) *MultiReadResponse {
	return &MultiReadResponse{
		response: response,
	}
}

func (m *MultiReadResponse) Context() Context {
	return m.response.Context()
}

func (m *MultiReadResponse) AddCookie(cookie *Cookie) {
	m.response.AddCookie(cookie)
}

func (m *MultiReadResponse) ContentLength() int {
	return m.response.ContentLength()
}

func (m *MultiReadResponse) SetContentLength(len int) {
	m.response.SetContentLength(len)
}

func (m *MultiReadResponse) CharacterEncoding() string {
	return m.response.CharacterEncoding()
}

func (m *MultiReadResponse) SetCharacterEncoding(charset string) {
	m.response.SetCharacterEncoding(charset)
}

func (m *MultiReadResponse) ContentType() string {
	return m.response.ContentType()
}

func (m *MultiReadResponse) SetContentType(contentType string) {
	m.response.SetContentType(contentType)
}

func (m *MultiReadResponse) AddHeader(name string, value string) {
	m.response.AddHeader(name, value)
}

func (m *MultiReadResponse) SetHeader(name string, value string) {
	m.response.SetHeader(name, value)
}

func (m *MultiReadResponse) DeleteHeader(name string) {
	m.response.DeleteHeader(name)
}

func (m *MultiReadResponse) Header(name string) (string, bool) {
	return m.response.Header(name)
}

func (m *MultiReadResponse) HeaderNames() []string {
	return m.HeaderNames()
}

func (m *MultiReadResponse) Headers(name string) []string {
	return m.response.Headers(name)
}

func (m *MultiReadResponse) Status() Status {
	return m.response.Status()
}

func (m *MultiReadResponse) SetStatus(status Status) {
	m.response.SetStatus(status)
}

func (m *MultiReadResponse) Writer() io.Writer {
	//TODO implement me
	panic("implement me")
}

func (m *MultiReadResponse) Flush() error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (m *MultiReadResponse) IsCommitted() bool {
	//TODO implement me
	panic("implement me")
}

func (m *MultiReadResponse) Reset() {
	//TODO implement me
	panic("implement me")
}

func (m *MultiReadResponse) CopyBodyToResponse() error {
	return nil
}
