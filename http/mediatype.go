package http

import (
	"mime"
	"strings"
)

type MediaType string

const (
	MediaTypeAll               MediaType = "*/*"
	MediaTypeAtomXml           MediaType = "application/atom+xml"
	MediaTypeCbor              MediaType = "application/cbor"
	MediaTypeFormUrlencoded    MediaType = "application/x-www-form-urlencoded"
	MediaTypeGraphqlResponse   MediaType = "application/graphql-response+json"
	MediaTypeJson              MediaType = "application/json"
	MediaTypeJsonUtf8          MediaType = "application/json;charset=UTF-8"
	MediaTypeYaml              MediaType = "application/yaml"
	MediaTypeYamlUtf8          MediaType = "application/yaml;charset=UTF-8"
	MediaTypeOctetStream       MediaType = "application/octet-stream"
	MediaTypePdf               MediaType = "application/pdf"
	MediaTypeProblemJson       MediaType = "application/problem+json"
	MediaTypeProblemJsonUtf8   MediaType = "application/problem+json;charset=UTF-8"
	MediaTypeProblemXml        MediaType = "application/problem+xml"
	MediaTypeProtobuf          MediaType = "application/x-protobuf"
	MediaTypeRssXml            MediaType = "application/rss+xml"
	MediaTypeStreamJson        MediaType = "application/stream+json"
	MediaTypeXhtmlXml          MediaType = "application/xhtml+xml"
	MediaTypeXml               MediaType = "application/xml"
	MediaTypeImageGif          MediaType = "image/gif"
	MediaTypeImageJpeg         MediaType = "image/jpeg"
	MediaTypeImagePng          MediaType = "image/png"
	MediaTypeMultipartFormData MediaType = "multipart/form-data"
	MediaTypeMultipartMixed    MediaType = "multipart/mixed"
	MediaTypeMultipartRelated  MediaType = "multipart/related"
	MediaTypeTextEventStream   MediaType = "text/event-stream"
	MediaTypeTextHtml          MediaType = "text/html"
	MediaTypeTextMarkdown      MediaType = "text/markdown"
	MediaTypeTextPlain         MediaType = "text/plain"
	MediaTypeTextXml           MediaType = "text/xml"
)

type cacheMediaType struct {
	typ     string
	subtype string
	charset string
	params  map[string]string
}

var (
	cacheMediaTypes = map[MediaType]cacheMediaType{}
)

func init() {
	parseMediaType(MediaTypeAll)
	parseMediaType(MediaTypeAtomXml)
	parseMediaType(MediaTypeCbor)
	parseMediaType(MediaTypeFormUrlencoded)
	parseMediaType(MediaTypeGraphqlResponse)
	parseMediaType(MediaTypeJson)
	parseMediaType(MediaTypeJsonUtf8)
	parseMediaType(MediaTypeYaml)
	parseMediaType(MediaTypeYamlUtf8)
	parseMediaType(MediaTypeOctetStream)
	parseMediaType(MediaTypePdf)
	parseMediaType(MediaTypeProblemJson)
	parseMediaType(MediaTypeProblemJsonUtf8)
	parseMediaType(MediaTypeProblemXml)
	parseMediaType(MediaTypeProtobuf)
	parseMediaType(MediaTypeRssXml)
	parseMediaType(MediaTypeStreamJson)
	parseMediaType(MediaTypeXhtmlXml)
	parseMediaType(MediaTypeXml)
	parseMediaType(MediaTypeImageGif)
	parseMediaType(MediaTypeImageJpeg)
	parseMediaType(MediaTypeImagePng)
	parseMediaType(MediaTypeMultipartFormData)
	parseMediaType(MediaTypeMultipartFormData)
	parseMediaType(MediaTypeMultipartMixed)
	parseMediaType(MediaTypeMultipartMixed)
	parseMediaType(MediaTypeMultipartRelated)
	parseMediaType(MediaTypeTextEventStream)
	parseMediaType(MediaTypeTextEventStream)
	parseMediaType(MediaTypeTextHtml)
	parseMediaType(MediaTypeTextMarkdown)
	parseMediaType(MediaTypeTextPlain)
	parseMediaType(MediaTypeTextXml)
}

func parseMediaType(mediaType MediaType) {
	fullType, params, _ := mime.ParseMediaType(string(mediaType))
	typ, subtype, _ := strings.Cut(fullType, "/")

	cache := cacheMediaType{
		typ:     typ,
		subtype: subtype,
		charset: "",
		params:  params,
	}

	charset, exists := params["charset"]
	if exists {
		cache.charset = charset
	}

	cacheMediaTypes[mediaType] = cache
}

func (m MediaType) Type() string {
	if val, ok := cacheMediaTypes[m]; ok {
		return val.typ
	}

	return ""
}

func (m MediaType) IsWildcardType() bool {
	if val, ok := cacheMediaTypes[m]; ok {
		return val.typ == "*"
	}

	return false
}

func (m MediaType) Subtype() string {
	if val, ok := cacheMediaTypes[m]; ok {
		return val.subtype
	}

	return ""
}

func (m MediaType) IsWildcardSubtype() bool {
	if val, ok := cacheMediaTypes[m]; ok {
		return val.subtype == "*"
	}

	return false
}

func (m MediaType) Charset() string {
	if val, ok := cacheMediaTypes[m]; ok {
		return val.charset
	}

	return ""
}

func (m MediaType) Parameter(name string) (string, bool) {
	if val, ok := cacheMediaTypes[m]; ok {
		param, exists := val.params[name]
		return param, exists
	}

	return "", false
}
