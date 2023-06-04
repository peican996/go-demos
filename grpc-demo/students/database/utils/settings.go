package utils

import (
	"log"

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
	relativePath := "database-demo/conf/config.ini"
	file, err := ini.Load(relativePath)
	if err != nil {
		log.Fatalf("配置文件读取错误，请检查文件路径: %v, ", err)
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
