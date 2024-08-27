package web

import (
	"codnect.io/procyon-web/http"
	"golang.org/x/exp/maps"
	"io"
	stdhttp "net/http"
	"net/url"
)

type defaultHttpServerRequest struct {
	req    *stdhttp.Request
	ctx    http.Context
	reader io.Reader

	queryCache   url.Values
	cookiesCache []*http.Cookie
}

func (r *defaultHttpServerRequest) WithReader(reader io.Reader) http.Request {
	if reader == nil {
		panic("nil reader")
	}

	copyRequest := new(defaultHttpServerRequest)
	*copyRequest = *r
	copyRequest.reader = reader
	return copyRequest
}

func (r *defaultHttpServerRequest) Context() http.Context {
	return r.ctx
}

func (r *defaultHttpServerRequest) initCookieCache() {
	if r.cookiesCache == nil {
		r.cookiesCache = parseCookies(r.req.Header)
	}
}

func (r *defaultHttpServerRequest) Cookie(name string) (*http.Cookie, bool) {
	r.initCookieCache()

	for _, cookie := range r.cookiesCache {
		if cookie.Name == name {
			return cookie, true
		}
	}

	return nil, false
}

func (r *defaultHttpServerRequest) Cookies() []*http.Cookie {
	r.initCookieCache()
	return r.cookiesCache
}

func (r *defaultHttpServerRequest) initQueryCache() {
	if r.queryCache == nil {
		if r.req != nil && r.req.URL != nil {
			r.queryCache = r.req.URL.Query()
		}
	}
}

func (r *defaultHttpServerRequest) QueryParameter(name string) (string, bool) {
	r.initQueryCache()

	values, ok := r.queryCache[name]
	if ok {
		return values[0], true
	}

	return "", false
}

func (r *defaultHttpServerRequest) QueryParameterNames() []string {
	r.initQueryCache()
	return maps.Keys(r.queryCache)
}

func (r *defaultHttpServerRequest) QueryParameters(name string) []string {
	r.initQueryCache()
	return r.queryCache[name]
}

func (r *defaultHttpServerRequest) QueryString() string {
	r.initQueryCache()
	return r.req.URL.RawQuery
}

func (r *defaultHttpServerRequest) Header(name string) (string, bool) {
	values := r.req.Header.Values(name)

	if len(values) != 0 {
		return values[0], true
	}

	return "", false
}

func (r *defaultHttpServerRequest) HeaderNames() []string {
	return maps.Keys(r.req.Header)
}

func (r *defaultHttpServerRequest) Headers(name string) []string {
	return r.req.Header.Values(name)
}

func (r *defaultHttpServerRequest) Path() string {
	return r.req.URL.Path
}

func (r *defaultHttpServerRequest) Method() http.Method {
	return http.Method(r.req.Method)
}

func (r *defaultHttpServerRequest) Reader() io.Reader {
	if r.reader != nil {
		return r.reader
	}

	return r.req.Body
}

func (r *defaultHttpServerRequest) Scheme() string {
	return ""
}

func (r *defaultHttpServerRequest) IsSecure() bool {
	return false
}
