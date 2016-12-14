package kiss

import (
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/unrolled/render"
	"gopkg.in/square/go-jose.v1/json"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (c *Context) URLParams() map[string]string {
	return mux.Vars(c.Request)
}

func (c *Context) URLParam(key string) string {
	return c.URLParams()[key]
}

func (c *Context) URLParamUint(key string) (uint, error) {
	v := c.URLParam(key)
	iv, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	return uint(iv), nil
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
