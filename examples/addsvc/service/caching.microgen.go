// Code generated by microgen 1.0.0-alpha. DO NOT EDIT.

package service

import (
	"context"
	service "github.com/devimteam/microgen/examples/addsvc"
	log "github.com/go-kit/kit/log"
)

// Cache interface uses for middleware as key-value storage for requests.
type Cache interface {
	Set(key, value interface{}) (err error)
	Get(key interface{}) (value interface{}, err error)
}

func CachingMiddleware(cache Cache) Middleware {
	return func(next service.Service) service.Service {
		return &cachingMiddleware{
			cache: cache,
			next:  next,
		}
	}
}

type cachingMiddleware struct {
	cache  Cache
	logger log.Logger
	next   service.Service
}

func (M cachingMiddleware) Sum(ctx context.Context, a int, b int) (res0 int, res1 error) {
	return M.next.Sum(ctx, a, b)
}

func (M cachingMiddleware) Concat(ctx context.Context, a string, b string) (res0 string, res1 error) {
	return M.next.Concat(ctx, a, b)
}

type sumResponseCacheEntity struct {
	Result int
}
type concatResponseCacheEntity struct {
	Result string
}
