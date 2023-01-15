package models

import (
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);comment:项目名称"`
	Desc    string
	StartAt null.Time
	EndAt   null.Time
	Owner   uint `gorm:"comment:项目拥有者"`
}
