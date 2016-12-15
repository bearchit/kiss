package kiss

import "net/http"

type handlerFunc func(*Context)

type handler struct {
	Handler handlerFunc
	Logger  *logger
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := Context{
		ResponseWriter: w,
		Request:        r,
		Logger:         h.Logger,
	}

	h.Handler(&c)
}
