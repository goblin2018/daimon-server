package dao

import (
	"daimon/models"
	"daimon/pkg/mysql"

	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao() *TaskDao {
	return &TaskDao{mysql.GetDB()}
}

func (d *TaskDao) AddTask(t *models.Task) error {
	last := new(models.Task)
	d.Model(&models.Task{}).Where("task_group_id = ?", t.TaskGroupId).Order("serial DESC").First(&last)
	if last.ID <= 0 {
		t.Serial = 0
	} else {
		t.Serial = (last.Serial/SerialGap + 1) * SerialGap
	}
	return d.Create(t).Error
}

func (d *TaskDao) DelTask(id uint) {
	d.Unscoped().Model(&models.Task{}).Where("id = ?", id).Delete(&models.Task{})
}

func (d *TaskDao) UpdateTask(t *models.Task) error {
	return d.Model(&models.Task{}).Where("id = ?", t.ID).Omit("id, deleted_at").Updates(t).Error
}

func (d *TaskDao) UpdateTaskInfo(t *models.Task) error {
	return d.Model(&models.Task{}).Where("id = ?", t.ID).Updates(map[string]interface{}{
		"name":     t.Name,
		"start_at": t.StartAt,
		"end_at":   t.EndAt,
	}).Error
}

func (d *TaskDao) UpdateTaskStatus(t *models.Task) error {
	return d.Model(&models.Task{}).Where("id = ?", t.ID).Update("status", t.Status).Error
}

func (d *TaskDao) GetTaskById(id uint) (t *models.Task, err error) {
	t = new(models.Task)
	err = d.Model(&models.Task{}).Where("id = ?", id).First(t).Error
	return
}

func (d *TaskDao) UpdateTaskByMove(t *models.Task) error {
	return d.Model(&models.Task{}).Where("id = ?", t.ID).Updates(map[string]interface{}{
		"serial":        t.Serial,
		"task_group_id": t.TaskGroupId,
	}).Error
}

func (d *TaskDao) GetTasksByTaskGroupId(taskGroupId uint) (tasks []*models.Task) {
	d.Model(&models.Task{}).Where("task_group_id = ?", taskGroupId).Find(&tasks)
	return
}
