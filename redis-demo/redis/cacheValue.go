package redis

import (
	"demos/redis-demo/model"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func init() {

}

func CacheValue(user model.User) {
	// 序列化为JSON字符串
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatal("序列化用户数据失败:", err)
	}
	// 设置到Redis缓存中
	err = Client.Set(fmt.Sprintf("user:%d", user.ID), string(userJSON), time.Hour).Err()
	if err != nil {
		log.Fatal("设置缓存失败:", err)
	}
}
