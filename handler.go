package kiss

import "net/http"

type HandlerFunc func(*Context)

type Handler struct {
	Handler HandlerFunc
}

func NewHandler(hf HandlerFunc) *Handler {
	return &Handler{hf}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := Context{
		ResponseWriter: w,
		Request:        r,
	}

	h.Handler(&c)
}
