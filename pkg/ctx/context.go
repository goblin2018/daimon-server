package ctx

import (
	"github.com/gin-gonic/gin"
)

// Custon Context data and funcs
type Context struct {
	*gin.Context
}

func (c *Context) SetUserID(id string) {
	c.Set("user_id", id)

}

func (c *Context) GetUserID() string {
	return c.GetString("user_id")
}

func (c *Context) SetPhone(phone string) {
	c.Set("phone", phone)
}

func (c *Context) GetPhone() string {
	return c.GetString("phone")

}
