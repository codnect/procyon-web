package web

import (
	"codnect.io/procyon-web/http"
	"golang.org/x/exp/maps"
	"io"
	stdhttp "net/http"
	"net/url"
	"strconv"
)

type defaultHttpServerResponse struct {
	responseWriter stdhttp.ResponseWriter
	ctx            http.Context
	writer         io.Writer

	headers        stdhttp.Header
	statusCode     http.Status
	writtenHeaders bool
	writerUsed     bool
}

func (r *defaultHttpServerResponse) WithWriter(writer io.Writer) http.Response {
	if writer == nil {
		panic("nil writer")
	}

	copyResponse := new(defaultHttpServerResponse)
	*copyResponse = *r
	copyResponse.writer = writer
	copyResponse.writerUsed = false

	return copyResponse
}

func (r *defaultHttpServerResponse) Context() http.Context {
	return r.ctx
}

func (r *defaultHttpServerResponse) AddCookie(cookie *http.Cookie) {
	if r.writtenHeaders {
		return
	}

	path := cookie.Path
	if path == "" {
		path = "/"
	}

	stdCookie := &stdhttp.Cookie{
		Name:     cookie.Name,
		Value:    url.QueryEscape(cookie.Value),
		Path:     path,
		Domain:   cookie.Domain,
		Expires:  cookie.Expires,
		MaxAge:   cookie.MaxAge,
		Secure:   cookie.Secure,
		HttpOnly: cookie.HttpOnly,
		SameSite: stdhttp.SameSite(cookie.SameSite),
	}

	if v := stdCookie.String(); v != "" {
		r.headers.Add("Set-Cookie", v)
	}
}

func (r *defaultHttpServerResponse) ContentLength() int {
	length := r.headers.Get("Content-Length")

	if length != "" {
		val, err := strconv.Atoi(length)
		if err != nil {
			return 0
		}

		return val
	}

	return 0
}

func (r *defaultHttpServerResponse) SetContentLength(len int) {
	if r.writtenHeaders {
		return
	}

	r.headers.Add("Content-Length", strconv.Itoa(len))
}

func (r *defaultHttpServerResponse) CharacterEncoding() string {
	return ""
}

func (r *defaultHttpServerResponse) SetCharacterEncoding(charset string) {

}

func (r *defaultHttpServerResponse) ContentType() string {
	return r.headers.Get("Content-Type")
}

func (r *defaultHttpServerResponse) SetContentType(contentType string) {
	if r.writtenHeaders {
		return
	}

	r.headers.Add("Content-Type", contentType)
}

func (r *defaultHttpServerResponse) AddHeader(name string, value string) {
	if r.writtenHeaders {
		return
	}

	r.headers.Add(name, value)
}

func (r *defaultHttpServerResponse) SetHeader(name string, value string) {
	if r.writtenHeaders {
		return
	}

	r.headers.Set(name, value)
}

func (r *defaultHttpServerResponse) DeleteHeader(name string) {
	if r.writtenHeaders {
		return
	}

	r.headers.Del(name)
}

func (r *defaultHttpServerResponse) Header(name string) (string, bool) {
	values := r.headers.Values(name)

	if len(values) != 0 {
		return values[0], true
	}

	return "", false
}

func (r *defaultHttpServerResponse) HeaderNames() []string {
	return maps.Keys(r.headers)
}

func (r *defaultHttpServerResponse) Headers(name string) []string {
	return r.headers.Values(name)
}

func (r *defaultHttpServerResponse) Status() http.Status {
	return r.statusCode
}

func (r *defaultHttpServerResponse) SetStatus(status http.Status) {
	if r.writtenHeaders {
		return
	}

	r.statusCode = status
}

func (r *defaultHttpServerResponse) Writer() io.Writer {
	r.writerUsed = true

	if r.writer != nil {
		return r.writer
	}

	return r.responseWriter
}

func (r *defaultHttpServerResponse) Flush() {
	r.writeHeaders()
	if !r.writerUsed {
		return
	}

	if r.writer == nil {
		r.responseWriter.WriteHeader(int(r.statusCode))
	}
	// flush data
}

func (r *defaultHttpServerResponse) IsCommitted() bool {
	return false
}

func (r *defaultHttpServerResponse) Reset() {
	r.headers = stdhttp.Header{}
}

func (r *defaultHttpServerResponse) writeHeaders() {
	if !r.writtenHeaders {
		for key, values := range r.headers {
			if len(values) == 1 {
				r.responseWriter.Header().Set(key, values[0])
			} else {
				for _, value := range values {
					r.responseWriter.Header().Add(key, value)
				}
			}
		}
		r.writtenHeaders = true
	}
}
