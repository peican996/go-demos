package main

import (
	"database/sql"
	"demos/gorm-demo/model"
	"demos/gorm-demo/utils"
	"time"
)

func main() {
	for i := 0; i < 2; i++ {
		user := model.User{Name: utils.GetFullName(), Email: utils.GetEmail(), Age: uint8(utils.GetAge()), Birthday: time.Now(), MemberNumber: string(utils.GetRandomInt()), ActivatedAt: sql.NullTime{Time: time.Now(), Valid: true}, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		model.CreateUser(&user)
		//fmt.Println(model.CreateUser(&user))
	}
}
