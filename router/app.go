package router

import (
	"daimon/pkg/ctx"
	"daimon/router/controllers"
)

func InitRouter() *ctx.Engine {

	app := ctx.Default()

	v1 := app.Group("/api/v1")
	ws := app.Group("/ws/v1")
	controllers.NewProjectController().RegisterRouters(v1)

	// Todo 这里是 websocket服务的路由
	controllers.NewWsController().RegisterRouters(ws)

	return app
}
