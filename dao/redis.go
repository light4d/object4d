package dao

import (
	"github.com/go-redis/redis"
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/object4d/common/server"
)

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     server.APPConfig.Redis.Addr,
		Password: server.APPConfig.Redis.Password,
		DB:       server.APPConfig.Redis.DB,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		log.Fatal(log.Fields{
			"redis":  pong,
			"err":    err,
			"detail": server.APPConfig.Redis,
		})

	}
	return client
}
