package main

import (
	"database/sql"
	"demos/gorm-demo/model"
	"fmt"
	"time"
)

func main() {
	user := model.User{Name: "Jinzhu", Email: "1414151@14187.com", Age: 18, Birthday: time.Now(), MemberNumber: "141", ActivatedAt: sql.NullTime{Time: time.Now(), Valid: true}, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	fmt.Println(model.CreateUser(&user))
}
