package controllers

import (
	"daimon/api"
	"daimon/pkg/ctx"
	"daimon/pkg/e"
	"daimon/services/task"
)

type TaskController struct {
	s *task.TaskService
}

func NewTaskController() *TaskController {
	return &TaskController{task.NewTaskService()}
}

func (co *TaskController) RegisterRouters(en *ctx.RouterGroup) {
	t := en.Group("/task")
	t.POST("", co.addTask)
	t.PUT("/status", co.updateTaskStatus)
	t.POST("/move", co.moveTask)
	t.PUT("", co.updateTask)
}

func (co *TaskController) addTask(c *ctx.Context) {
	req := new(api.Task)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}

	res, err := co.s.AddTask(c, req)
	c.JSON(res, err)
}

func (co *TaskController) updateTaskStatus(c *ctx.Context) {
	req := new(api.Task)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}

	res, err := co.s.UpdateTaskStatus(c, req)
	c.JSON(res, err)
}

func (co *TaskController) moveTask(c *ctx.Context) {
	req := new(api.MoveTaskReq)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}

	res, err := co.s.MoveTask(c, req)
	c.JSON(res, err)
}

func (co *TaskController) updateTask(c *ctx.Context) {
	req := new(api.Task)
	if err := c.ShouldBind(req); err != nil {
		c.Fail(e.InvalidParams.Add(err.Error()))
		return
	}

	res, err := co.s.UpdateTask(c, req)
	c.JSON(res, err)
}
