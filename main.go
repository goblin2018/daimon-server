package main

import (
	"daimon/pkg/conf"
	"daimon/pkg/log"
	"daimon/router"
)

func main() {
	log.InitLogger()
	// redis.Init()
	// mysql.Init()

	app := router.InitRouter()

	ac := conf.C.App
	app.Run(ac.Host + ":" + ac.Port)

}
