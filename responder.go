package kiss

import "net/http"

func (c *Context) Respond(status int, v interface{}) {
	if v == nil {
		c.JSON(status, map[string]string{
			"message": http.StatusText(status),
		})
	} else {
		c.JSON(status, v)
	}
}

func (c *Context) NotImpl() {
	c.Respond(http.StatusInternalServerError, map[string]string{
		"message": "Not implemented yet",
	})
}

func (c *Context) JustOK() {
	c.Respond(http.StatusOK, nil)
}

func (c *Context) OK(v interface{}) {
	c.Respond(http.StatusOK, v)
}

func (c *Context) BadRequest(v interface{}) {
	c.Respond(http.StatusBadRequest, v)
}

func (c *Context) Unauthorized(v interface{}) {
	c.Respond(http.StatusUnauthorized, v)
}

func (c *Context) Forbidden(v interface{}) {
	c.Respond(http.StatusForbidden, v)
}

func (c *Context) NotFound(v interface{}) {
	c.Respond(http.StatusNotFound, v)
}

func (c *Context) InternalServerError(v interface{}) {
	c.Respond(http.StatusInternalServerError, v)
}
