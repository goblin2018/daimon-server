package api

type User struct {
	Id     uint   `json:"id" form:"id"`
	UserId string `json:"userId" form:"userId"`
	Phone  string `json:"phone" form:"phone"`
	Name   string `json:"name" form:"name"`
	Duty   string `json:"duty" form:"duty"`
	Status int    `json:"status" form:"status"`
}
