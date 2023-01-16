package api

type Project struct {
	Id   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Desc string `json:"desc" form:"desc"`
}
