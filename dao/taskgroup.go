package dao

import (
	"daimon/models"
	"daimon/pkg/mysql"

	"gorm.io/gorm"
)

const SerialGap = 65536

type TaskGroupDao struct {
	*gorm.DB
}

func NewTaskGroupDao() *TaskGroupDao {
	return &TaskGroupDao{mysql.GetDB()}
}

func (d *TaskGroupDao) AddTaskGroup(t *models.TaskGroup) error {
	// 看下上一个最大值是多少
	last := new(models.TaskGroup)
	d.Model(&models.TaskGroup{}).Where("project_id = ?", t.ProjectId).Order("serial DESC").First(&last)
	if last.ID <= 0 {
		t.Serial = 0
	} else {
		t.Serial = (last.Serial/SerialGap + 1) * SerialGap
	}

	return d.Create(t).Error
}

func (d *TaskGroupDao) InsertToFirst(id, projectId uint) (serial int) {
	// 看下上一个最大值是多少
	first := new(models.TaskGroup)
	d.Model(&models.TaskGroup{}).Where("project_id = ?", projectId).Order("serial ASC").First(&first)
	serial = (first.Serial/SerialGap - 1) * SerialGap
	d.Model(&models.TaskGroup{}).Where("id = ?", id).Update("serial", serial)
	return
}

func (d *TaskGroupDao) InsertToLast(id, projectId uint) (serial int) {
	// 看下上一个最大值是多少
	last := new(models.TaskGroup)
	d.Model(&models.TaskGroup{}).Where("project_id = ?", projectId).Order("serial DESC").First(&last)
	serial = (last.Serial/SerialGap + 1) * SerialGap
	d.Model(&models.TaskGroup{}).Where("id = ?", id).Update("serial", serial)
	return
}

func (d *TaskGroupDao) DelTaskGroup(id uint) {
	d.Unscoped().Model(&models.Task{}).Where("task_group_id = ?", id).Delete(&models.Task{})
	d.Unscoped().Model(&models.TaskGroup{}).Where("id = ?", id).Delete(&models.TaskGroup{})
}

func (d *TaskGroupDao) UpdateTaskGroup(t *models.TaskGroup) error {
	return d.Model(&models.TaskGroup{}).Where("id = ?", t.ID).Updates(t).Error
}

func (d *TaskGroupDao) UpdateSerial(id uint, serial int) error {
	return d.Model(&models.TaskGroup{}).Where("id = ?", id).Update("serial", serial).Error
}

func (d *TaskGroupDao) GetTaskGroupSerialById(id uint) (serial int) {
	d.Model(&models.TaskGroup{}).Select("serial").Where("id = ?", id).First(&serial)
	return
}

func (d *TaskGroupDao) GetSerialsByProjectId(id uint) (tgs []*models.TaskGroup) {
	d.Model(&models.TaskGroup{}).Select("id, serial").Where("project_id = ?", id).Order("serial ASC").Find(&tgs)
	return
}
