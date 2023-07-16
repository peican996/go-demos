package common

import (
	"bufio"
	"demos/weibospider-demo/config/model"
	"fmt"
	"os"
)

var WorkPath string

var Info = new(StartInfo)

type StartInfo struct {
	Config     model.Config
	UserIds    []string
	ScriptPath string
}

func InitConfig() {
	WorkPath, _ = os.Getwd()
	Info = &StartInfo{
		Config:     model.ConfigJson,
		UserIds:    GetUsersId(),
		ScriptPath: "D:\\code\\python\\weiboSpider\\",
	}
}

func GetUsersId() []string {
	// 打开文件
	file, err := os.Open(WorkPath + "\\weibospider-demo\\conf\\userId.conf")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return nil
	}
	defer file.Close()
	// 创建一个字符串切片来存储每一行的内容
	var lines []string
	// 创建一个 Scanner 并将文件设置为输入源
	scanner := bufio.NewScanner(file)
	// 逐行读取文件
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件错误:", err)
		return nil
	}
	// 打印每一行的内容
	for _, line := range lines {
		fmt.Println(line)
	}
	return lines
}
