package kiss

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/unrolled/render"
)

func (c *Context) UrlParams() map[string]string {
	return mux.Vars(c.Request)
}

func (c *Context) UrlParam(key string) string {
	return strings.TrimSpace(c.UrlParams()[key])
}

func (c *Context) UrlParamInt(key string) int {
	v := c.UrlParam(key)
	iv, _ := strconv.Atoi(v)
	return iv
}

func (c *Context) UrlParamUint(key string) uint {
	return uint(c.UrlParamInt(key))
}

func (c *Context) QueryParams() url.Values {
	return c.Request.URL.Query()
}

func (c *Context) QueryParam(key string) string {
	values := c.QueryParams()[key]
	if len(values) > 0 {
		return values[0]
	}

	return ""
}

func (c *Context) QueryParamValues(key string) []string {
	return c.QueryParams()[key]
}

func (c *Context) QueryParamInt(key string) int {
	v := c.QueryParam(key)
	iv, _ := strconv.Atoi(v)
	return iv
}

func (c *Context) QueryParamUint(key string) uint {
	return uint(c.QueryParamInt(key))
}

func (c *Context) JSON(status int, v interface{}) {
	render.New().JSON(c.ResponseWriter, status, v)
}

func (c *Context) BindForm(v interface{}) error {
	if err := c.Request.ParseForm(); err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	if err := decoder.Decode(v, c.Request.Form); err != nil {
		return err
	}

	return nil
}

func (c *Context) ValidateForm(v interface{}) error {
	if err := c.BindForm(v); err != nil {
		return err
	}

	if _, err := govalidator.ValidateStruct(v); err != nil {
		return err
	}

	return nil
}
