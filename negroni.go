package kiss

import "net/http"

type NegroniMiddleware struct {
	Middleware Middleware
}

func NewNegroniMiddleware(m Middleware) *NegroniMiddleware {
	return &NegroniMiddleware{m}
}

func (m *NegroniMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	c := Context{
		ResponseWriter: w,
		Request:        r,
	}

	m.Middleware.ServeHTTP(&c, next)
}
