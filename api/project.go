package api

import (
	"gopkg.in/guregu/null.v4"
)

type Project struct {
	Id         uint         `json:"id" form:"id"`
	Name       string       `json:"name" form:"name"`
	StartAt    null.Time    `json:"startAt" form:"startAt"`
	EndAt      null.Time    `json:"endAt" form:"endAt"`
	Desc       string       `json:"desc" form:"desc"`
	TaskGroups []*TaskGroup `json:"taskGroups" form:"taskGroups"`
}
