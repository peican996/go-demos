package redis

import (
	"log"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr: "172.22.0.2:6379",
	})
	_, err := Client.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
}
