package kiss

import (
	"net/http"

	"github.com/mangoplate/kiss/log"
)

type handlerFunc func(*Context)

type handler struct {
	Handler handlerFunc
	Logger  *log.Logger
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := Context{
		ResponseWriter: w,
		Request:        r,
		Logger:         h.Logger,
	}

	h.Handler(&c)
}
