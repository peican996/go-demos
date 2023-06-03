package main

import (
	"demos/database-demo/database"
	"demos/database-demo/utils"
	"fmt"
	"log"
)

func main() {
	db, err := database.NewDBManager(utils.GetDatabaseURL())
	if err != nil {
		log.Fatalln("获取数据连接失败!!!!!")
	}

	rows, err := db.Query("select * from students")
	if err != nil {
		log.Println("查询数据失败!!!!!")
	}

	// 处理查询结果
	for rows.Next() {
		var id int
		var name string
		var age int
		var gender string
		var studentNumber string

		err := rows.Scan(&id, &name, &age, &gender, &studentNumber)
		if err != nil {
			log.Fatal(err.Error())
		}

		// 处理每一行的数据
		fmt.Println(id)
	}
}
