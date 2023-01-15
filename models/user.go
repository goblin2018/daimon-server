package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `gorm:"type:varchar(20);index:name;comment:姓名"`
	Phone  string `gorm:"type:char(11);uniqueIndex:phone;comment:电话号码"`
	Duty   string `gorm:"type:varchar(32)"`
	Status int
}
