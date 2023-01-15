package api

type TaskGroup struct {
	Id        uint    `json:"id" form:"id"`
	ProjectId uint    `json:"projectId" form:"projectId"`
	Name      string  `json:"name" form:"name"`
	Serial    int     `json:"serial" form:"serial"`
	Tasks     []*Task `json:"tasks" form:"tasks"`
}

type MoveTaskGroupResp struct {
	Serial     int   `json:"serial" form:"serial"`
	UpdateAll  bool  `json:"updateAll" form:"updateAll"`
	AllSerials []int `json:"allSerials" form:"allSerials"`
}

type MoveTaskGroupReq struct {
	ProjectId uint `json:"projectId" form:"projectId"`
	Id        uint `json:"id" form:"id"`     // 要移动的id
	Prev      uint `json:"prev" form:"prev"` // 移动之后前面的id
	Next      uint `json:"next" form:"next"` // 移动之后后面的id
}
