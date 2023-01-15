package api

import "gopkg.in/guregu/null.v4"

type Task struct {
	Id          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Status      int       `json:"status" form:"status"`
	Owner       uint      `json:"owner" form:"owner"`
	PStartAt    null.Time `json:"pStartAt" form:"pStartAt"`
	PEndAt      null.Time `json:"pEndAt" form:"pEndAt"`
	StartAt     null.Time `json:"startAt" form:"startAt"`
	EndAt       null.Time `json:"endAt" form:"endAt"`
	ProjectId   uint      `json:"projectId" form:"projectId"`
	TaskGroupId uint      `json:"taskGroupId" form:"taskGroupId"`
	Serial      int       `json:"serial" form:"serial"`
}

type MoveTaskReq struct {
	Id          uint `json:"id" form:"id"`
	Prev        uint `json:"prev" form:"prev"`
	Next        uint `json:"next" form:"next"`
	TaskGroupId uint `json:"taskGroupId" form:"taskGroupId"`
}

type MoveTaskResp struct {
}
