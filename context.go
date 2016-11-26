package kiss

import (
	"net/http"

	"strconv"

	"strings"

	"net/url"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

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

//func (c *Context) BindJSON(v interface{}) error {
//	decoder := json.NewDecoder(c.Request.Body)
//	if err := decoder.Decode(&v); err != nil {
//		return err
//	}
//	defer c.Request.Body.Close()
//
//	_, err := govalidator.ValidateStruct(&v)
//	return err
//}
