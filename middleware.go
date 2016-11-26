package kiss

import "net/http"

type MiddlewareFunc func(*Context, http.HandlerFunc)

type Middleware interface {
	ServeHTTP(*Context, http.HandlerFunc)
}
