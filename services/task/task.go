package task

import (
	"daimon/api"
	"daimon/dao"
	"daimon/models"
	"daimon/pkg/ctx"
	"daimon/pkg/utils"
)

type TaskService struct {
	dao *dao.TaskDao
}

func NewTaskService() *TaskService {
	return &TaskService{dao.NewTaskDao()}
}

func (s *TaskService) AddTask(c *ctx.Context, req *api.Task) (resp *api.Task, err error) {
	resp = new(api.Task)
	t := &models.Task{
		ProjectId:   req.ProjectId,
		Name:        req.Name,
		TaskGroupId: req.TaskGroupId,
	}
	err = s.dao.AddTask(t)
	resp.Id = t.ID
	return
}

func (s *TaskService) UpdateTaskStatus(c *ctx.Context, req *api.Task) (resp *api.Task, err error) {
	resp = new(api.Task)
	t := &models.Task{
		Status: req.Status,
	}
	t.ID = req.Id
	err = s.dao.UpdateTaskStatus(t)
	return
}

func (s *TaskService) MoveTask(c *ctx.Context, req *api.MoveTaskReq) (resp *api.MoveTaskResp, err error) {
	resp = new(api.MoveTaskResp)

	// 获取task信息
	task, _ := s.dao.GetTaskById(req.Id)
	task.TaskGroupId = req.TaskGroupId

	var prevTask, nextTask *models.Task
	// 获取 prev next 信息
	if req.Prev > 0 {
		prevTask, _ = s.dao.GetTaskById(req.Prev)
	}
	if req.Next > 0 {
		nextTask, _ = s.dao.GetTaskById(req.Next)
	}
	// 计算新的task serial

	// 新分组 没有任务
	if req.Prev == 0 && req.Next == 0 {
		task.Serial = 0
		s.dao.UpdateTaskByMove(task)
		return
	}

	// 新分组 向前插入
	if req.Prev == 0 && req.Next > 0 {
		// 向前插入
		task.Serial = (nextTask.Serial/dao.SerialGap - 1) * dao.SerialGap
		s.dao.UpdateTaskByMove(task)
		return
	}

	// 新分组 向后插入
	if req.Prev > 0 && req.Next == 0 {
		task.Serial = (prevTask.Serial/dao.SerialGap + 1) * dao.SerialGap
		s.dao.UpdateTaskByMove(task)
		return
	}
	// 中间插入
	// 间距足够
	if nextTask.Serial-prevTask.Serial >= 2 {
		task.Serial = (nextTask.Serial + prevTask.Serial) / 2
		s.dao.UpdateTaskByMove(task)
		return
	}
	// 间距足够 重排
	tasks := s.dao.GetTasksByTaskGroupId(req.TaskGroupId)
	idx := findTaskIndex(tasks, req.Id)

	// 没有更换分组
	if idx >= 0 {
		// 先删掉原有的元素
		tasks = utils.DeleteItem(tasks, idx)
	}
	// 插入元素
	insertIdx := findTaskIndex(tasks, req.Next)
	tasks = utils.InsertItem(tasks, insertIdx, task)

	// 更新序列号
	for i, t := range tasks {
		t.Serial = i * dao.SerialGap
		s.dao.UpdateTaskByMove(t)
	}

	return

}

func findTaskIndex(tgs []*models.Task, id uint) int {
	for idx, tg := range tgs {
		if tg.ID == id {
			return idx
		}
	}
	return -1
}

func (s *TaskService) UpdateTask(c *ctx.Context, req *api.Task) (resp *api.Task, err error) {
	resp = new(api.Task)

	t := &models.Task{
		Name:    req.Name,
		StartAt: req.StartAt,
		EndAt:   req.EndAt,
	}
	t.ID = req.Id
	s.dao.UpdateTaskInfo(t)
	return
}
