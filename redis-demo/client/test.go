package main

import (
	"demos/redis-demo/model"
	"demos/redis-demo/redis"
	"encoding/json"
	"fmt"
)

func main() {
	str := redis.Client.Get("user:218")
	if str.Val() != "" {
		// 解析JSON数据到结构体
		var user model.User
		err := json.Unmarshal([]byte(str.Val()), &user)
		if err != nil {
			fmt.Println("解析JSON失败:", err)
		}
		return
	}
	user := model.GetUserByID("218")
	redis.CacheValue(user)
}
