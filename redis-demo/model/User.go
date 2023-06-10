package model

import (
	"database/sql"
	"demos/redis-demo/database"
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

func GetUserByID(id string) User {
	var user User
	result := database.DB.Table("users").First(&user, id)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
	return user
}
