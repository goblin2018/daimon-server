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
	controllers.NewWsController().RegisterRouters(ws)
	// NewTaskGroupController().RegisterRouters(v1)
	// NewTaskController().RegisterRouters(v1)

	return app
}
