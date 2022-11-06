package client

import (
	"net/http"
	"time"
)

var DefaultTimeout = 500 * time.Millisecond

var DefaultConnectTimeout = 1500 * time.Millisecond

type ContentType int

const (
	// JSON represents a JSON Content Type
	JSON ContentType = iota

	// XML represents an XML Content Type
	XML

	// BYTES represents a plain Content Type
	BYTES
)

type (
	Transport interface {
		Get(url string) *Response
		Post(url string, body interface{}) *Response
		Put(url string, body interface{}) *Response
		Patch(url string, body interface{}) *Response
		Delete(url string) *Response
		Head(url string) *Response
		Options(url string) *Response
	}
	RequestBuilder struct {

		// Headers to be send in the request
		Headers http.Header

		//Connection timeout, it bounds the time spent obtaining a successful connection
		ConnectTimeout time.Duration

		// Base URL to be used for each Request. The final URL will be BaseURL + URL.
		BaseURL string

		// ContentType
		ContentType ContentType

		// Disable timeout and default timeout = no timeout
		DisableTimeout bool

		// Public for custom fine tuning
		Client *http.Client

		LogTime bool
	}
)

func New(header http.Header, baseURL string, connectTimeout time.Duration, logTime bool) Transport {
	return &RequestBuilder{
		Headers:        header,
		BaseURL:        baseURL,
		ConnectTimeout: connectTimeout,
		LogTime:        logTime,
	}
}

func (rb *RequestBuilder) Get(url string) *Response {
	return rb.doRequest(http.MethodGet, url, nil)
}

func (rb *RequestBuilder) Post(url string, body interface{}) *Response {
	return rb.doRequest(http.MethodPost, url, body)
}

func (rb *RequestBuilder) Put(url string, body interface{}) *Response {
	return rb.doRequest(http.MethodPut, url, body)
}

func (rb *RequestBuilder) Patch(url string, body interface{}) *Response {
	return rb.doRequest(http.MethodPatch, url, body)
}

func (rb *RequestBuilder) Delete(url string) *Response {
	return rb.doRequest(http.MethodDelete, url, nil)
}

func (rb *RequestBuilder) Head(url string) *Response {
	return rb.doRequest(http.MethodHead, url, nil)
}

func (rb *RequestBuilder) Options(url string) *Response {
	return rb.doRequest(http.MethodOptions, url, nil)
}
