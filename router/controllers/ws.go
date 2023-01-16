package controllers

import (
	"daimon/api"
	"daimon/pkg/ctx"
	"daimon/pkg/e"
	"daimon/services/websocket"
)

type WsController struct {
	s *websocket.WsService
}

func NewWsController() *WsController {
	return &WsController{s: websocket.NewService()}
}

func (co WsController) RegisterRouters(en *ctx.RouterGroup) {
	ws := en.Group("/")
	ws.GET("/connect", co.connect)
	ws.GET("/sendtome", co.sendToClient)
}

func (co *WsController) connect(c *ctx.Context) {
	co.s.Connect(c)
}

func (co *WsController) sendToClient(c *ctx.Context) {
	req := new(api.Message)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams)
		return
	}

	co.s.SendToMe(c, req)
	c.JSON(nil, nil)
}
