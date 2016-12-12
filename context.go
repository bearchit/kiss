package kiss

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/unrolled/render"
	"gopkg.in/square/go-jose.v1/json"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (c *Context) BindForm(decoder *schema.Decoder, v interface{}) error {
	if err := c.Request.ParseForm(); err != nil {
		return err
	}

	if err := decoder.Decode(v, c.Request.Form); err != nil {
		return err
	}

	return nil
}

func (c *Context) BindJSON(v interface{}) error {
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}
	defer c.Request.Body.Close()

	return nil
}

func (c *Context) JSON(status int, v interface{}) {
	render.New().JSON(c.ResponseWriter, status, v)
}
