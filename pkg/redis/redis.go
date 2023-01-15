package redis

import (
	"daimon/pkg/conf"

	gredis "github.com/go-redis/redis"
)

var RDB *gredis.Client

func Init() {
	config := conf.C.Redis
	//rbd path
	rdbAddr := config.Host + ":" + config.Port

	//client
	RDB = gredis.NewClient(&gredis.Options{
		Addr: rdbAddr,
		DB:   0, // use default DB
	})

	//ping
	_, err := RDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
