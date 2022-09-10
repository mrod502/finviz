package client

import (
	"bytes"
	"io"
	"net/http"
	"sync/atomic"
	"time"

	gocache "github.com/mrod502/go-cache"
)

func NewHttp[T any](parser Parser[T]) *Http[T] {
	return &Http[T]{
		Client:   http.DefaultClient,
		parse:    parser,
		useCache: &atomic.Bool{},
	}
}

type Parser[T any] func(io.Reader, *T) error

func SetHeaders(r *http.Request) {
	r.Header.Set("accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`)
	r.Header.Set("user-agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36`)

}

type Http[T any] struct {
	*http.Client
	parse    Parser[T]
	useCache *atomic.Bool
	cache    *gocache.Cache[*T, string]
}

func (h *Http[T]) Get(path string) (v *T, err error) {
	return h.withCacheAttempt(path, func() (*T, error) {
		return h.getBase(path)
	})
}

func (h *Http[T]) getBase(path string) (v *T, err error) {
	v = new(T)
	req, _ := buildRequest(path, "GET", nil)
	res, err := h.Do(req)
	if err != nil {
		return nil, err
	}
	err = h.parse(res.Body, v)
	return
}

func (h *Http[T]) WithCache(lifetime time.Duration) *Http[T] {
	if lifetime == 0 {
		return h
	}
	h.useCache.Store(true)
	h.cache = gocache.New[*T, string]().WithExpiration(lifetime)
	return h
}

func (h *Http[T]) NoCache() *Http[T] {
	h.useCache.Store(false)
	h.cache = nil
	return h
}

func (h *Http[T]) withCacheAttempt(key string, f func() (*T, error)) (*T, error) {
	if h.useCache.Load() {
		if h.cache.Exists(key) {
			return h.cache.Get(key)
		}
	}
	return f()
}

func buildRequest(uri, method string, body []byte) (*http.Request, error) {
	var bod io.Reader
	if body != nil {
		bod = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, uri, bod)
	if err != nil {
		return nil, err
	}
	SetHeaders(req)
	return req, nil
}
