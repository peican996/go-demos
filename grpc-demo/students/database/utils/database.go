package utils

import "fmt"

func GetDatabaseURL() string {
	Init()
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbName, DbPassWord, DbHost, DbPort, DbUser)
}
