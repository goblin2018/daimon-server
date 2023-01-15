package models

import "gorm.io/gorm"

type TaskGroup struct {
	gorm.Model
	Name      string `gorm:"type:varchar(200);comment:任务组名称"`
	ProjectId uint   `gorm:"comment:项目Id"`
	Serial    int    `gorm:"comment:任务排序编号"`
}
