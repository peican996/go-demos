package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("database-demo/conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadDBData(file)
}

func LoadDBData(file *ini.File) {
	DbName = file.Section("database").Key("username").MustString("root")
	DbPassWord = file.Section("database").Key("password").MustString("123456")
	DbHost = file.Section("database").Key("host").MustString("localhost")
	DbPort = file.Section("database").Key("port").MustString("3306")
	DbUser = file.Section("database").Key("database").MustString("test")
}
