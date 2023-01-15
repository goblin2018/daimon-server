package controllers

import (
	"daimon/api"
	"daimon/pkg/ctx"
	"daimon/pkg/e"
)

type UserController struct {
}

func (co UserController) RegisterRouters(en *ctx.RouterGroup) {
	u := en.Group("/user")
	u.POST("/login", co.login)
}

func (co UserController) login(c *ctx.Context) {

}

func (co *UserController) register(c *ctx.Context) {
	req := new(api.User)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}

	// res, err := c.
}
