package main

import (
	"demos/redis-demo/model"
	"demos/redis-demo/redis"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		str := redis.Client.Get("user:217")
		if str.Val() != "" {
			// 解析JSON数据到结构体
			var user model.User
			err := json.Unmarshal([]byte(str.Val()), &user)
			if err != nil {
				fmt.Println("解析JSON失败:", err)
			}
		}
	}
	cost := time.Since(start)
	fmt.Println("cost=", cost)
	start1 := time.Now()
	for i := 0; i < 10000; i++ {
		model.GetUserByID("218")
	}
	cost1 := time.Since(start1)
	fmt.Println("cost=", cost1)
	//redis.CacheValue(user)
}
