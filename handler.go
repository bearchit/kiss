package kiss

import (
	"net/http"
)

type Handler struct {
	*Resource
	HandlerFunc func(*Context)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.HandlerFunc(NewContext(
		h.Resource,
		w,
		r,
	))
}
