package server

import (
	"codnect.io/procyon-web/http"
	"io"
	stdhttp "net/http"
)

type Response struct {
	responseWriter stdhttp.ResponseWriter
	ctx            http.Context
	writer         io.Writer
}

func (r *Response) WithWriter(writer io.Writer) http.Response {
	if writer == nil {
		panic("nil writer")
	}

	copyResponse := new(Response)
	*copyResponse = *r
	copyResponse.writer = writer
	return copyResponse
}

func (r *Response) Context() http.Context {
	return r.ctx
}

func (r *Response) AddCookie(cookie *http.Cookie) {

}

func (r *Response) ContentLength() int {
	return 0
}

func (r *Response) SetContentLength(len int) {

}

func (r *Response) CharacterEncoding() string {
	return ""
}

func (r *Response) SetCharacterEncoding(charset string) {

}

func (r *Response) ContentType() string {
	return ""
}

func (r *Response) SetContentType(contentType string) {

}

func (r *Response) AddHeader(name string, value string) {

}

func (r *Response) SetHeader(name string, value string) {

}

func (r *Response) DeleteHeader(name string) {

}

func (r *Response) Header(name string) string {
	return ""
}

func (r *Response) HeaderNames() []string {
	return nil
}

func (r *Response) Headers(name string) []string {
	return nil
}

func (r *Response) Status() http.Status {
	return 0
}

func (r *Response) SetStatus(status http.Status) {

}

func (r *Response) Writer() io.Writer {
	return nil
}

func (r *Response) Flush() {

}

func (r *Response) IsCommitted() bool {
	return false
}

func (r *Response) Reset() {

}
