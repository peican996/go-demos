package redis

import (
	"database/sql"
	"demos/redis-demo/model"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CacheAllData() {
	// 创建MySQL数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(172.22.0.2:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(db)
	// 从MySQL中查询数据
	rows, err := db.Query("SELECT ID, Name, Email, Age, MemberNumber FROM users")
	if err != nil {
		log.Fatal("查询MySQL数据失败:", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(rows)
	// 遍历查询结果，设置到Redis缓存中
	for rows.Next() {
		var user model.User
		if err != nil {
			log.Fatal("解析日期失败:", err)
		}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.MemberNumber)
		if err != nil {
			log.Fatal("读取查询结果失败:", err)
		}
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
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("数据已设置到Redis缓存中")
}
