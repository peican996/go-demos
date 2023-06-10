package model

import (
	"database/sql"
	databaseUtils "demos/gorm-demo/utils"
	"log"
	"time"
)

type User struct {
	ID           uint         `gorm:"column:id; primary_key; auto_increment"`
	Name         string       `gorm:"column:name; size:255"`
	Email        string       `gorm:"column:email; size:255; unique"`
	Age          uint8        `gorm:"column:age"`
	Birthday     time.Time    `gorm:"column:birthday"`
	MemberNumber string       `gorm:"column:MemberNumber"`
	ActivatedAt  sql.NullTime `gorm:"column:ActivatedAt"`
	CreatedAt    time.Time    `gorm:"column:CreatedAt"`
	UpdatedAt    time.Time    `gorm:"column:UpdatedAt"`
}

func init() {
	log.Println("初始化数据库连接！！！")
	databaseUtils.InitDb()
}

func CreateUser(data *User) int {
	db := databaseUtils.DB.Table("users").Create(&data)
	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}
	return int(db.RowsAffected)
}
