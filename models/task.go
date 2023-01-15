package models

import (
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);comment:任务名称"`
	Status      int       `gorm:"comment:0-未开始,1-进行中,2-已完成"`
	Owner       uint      `gorm:"comment:任务负责人"`
	PStartAt    null.Time `gorm:"comment:计划开始时间"`
	PEndAt      null.Time `gorm:"comment:计划结束时间"`
	StartAt     null.Time `gorm:"comment:开始时间"`
	EndAt       null.Time `gorm:"comment:结束时间"`
	ProjectId   uint      `gorm:"comment:项目Id"`
	TaskGroupId uint      `gorm:"comment:任务组Id"`
	Serial      int       `gorm:"comment:任务排序编号"`
}
