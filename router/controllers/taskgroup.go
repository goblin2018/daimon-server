package controllers

import (
	"daimon/api"
	"daimon/pkg/ctx"
	"daimon/pkg/e"
	"daimon/services/task"
)

type TaskGroupController struct {
	s *task.TaskGroupService
}

func NewTaskGroupController() *TaskGroupController {
	return &TaskGroupController{s: task.NewTaskGroupService()}
}

func (co *TaskGroupController) RegisterRouters(en *ctx.RouterGroup) {
	tg := en.Group("/taskgroup")
	tg.POST("", co.addTaskGroup)
	tg.POST("/move", co.moveTaskGroup)
}

func (co *TaskGroupController) addTaskGroup(c *ctx.Context) {
	req := new(api.TaskGroup)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}

	res, err := co.s.AddTaskGroup(c, req)
	c.JSON(res, err)
}
func (co *TaskGroupController) moveTaskGroup(c *ctx.Context) {
	req := new(api.MoveTaskGroupReq)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}

	res, err := co.s.MoveTaskGroup(c, req)
	c.JSON(res, err)
}
