package main

import (
	"daimon/models"
	"daimon/pkg/log"
	"daimon/pkg/mysql"
)

func main() {

	log.InitLogger()
	mysql.Init()
	db := mysql.GetDB()
	var err error

	err = db.AutoMigrate(&models.Project{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.TaskGroup{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
