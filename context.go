package kiss

import (
	"net/http"

	"github.com/unrolled/render"
)

type Context struct {
	*Resource

	render  *render.Render
	Writer  http.ResponseWriter
	Request *http.Request
}

func NewContext(res *Resource, w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Resource: res,
		render:   render.New(),
		Writer:   w,
		Request:  r,
	}
}

func (c *Context) JSON(status int, v interface{}) error {
	return c.render.JSON(c.Writer, status, v)
}
