package utils

import (
	"demos/weibospider-demo/config/model"
	"encoding/json"
	"fmt"
	"os"
)

func RenewConfigJson(usrId string, dataStr string) {
	model.ConfigJson = model.Config{
		UserIDList:          []string{},
		Filter:              1,
		SinceDate:           "2018-01-01",
		EndDate:             "2023-05-15 18:59",
		RandomWaitPages:     []int{1, 5},
		RandomWaitSeconds:   []int{6, 10},
		GlobalWait:          [][]int{{1000, 3600}, {500, 2000}},
		WriteMode:           []string{"txt"},
		PicDownload:         0,
		VideoDownload:       0,
		FileDownloadTimeout: []int{5, 5, 10},
		ResultDirName:       0,
		Cookie:              "-D52JvpLOsVIh-w.; SUB=_2A25JtgQsDeRhGeNO6VoQ9C7IyjmIHXVrWKxkrDV6PUJbktANLUmjkW1NTxjllnUwltMgZLuswXNR34OvWTpVlrDc; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9Wh57jVs4GfJR_d7uc2_yGeB5JpX5K-hUgL.Fo-7eonpSh5XeK-2dJLoIEBLxKqL1hBLB.qLxK.L1-2L12qLxK-LB.qL1heLxKnLBK2L12et; SSOLoginState=1689416828; ALF=1692008828; XSRF-TOKEN=875dd7; WEIBOCN_FROM=1110006030; mweibo_short_token=01b225c94f; MLOGIN=1; M_WEIBOCN_PARAMS=luicode%3D20000174%26uicode%3D20000174",
		MysqlConfig: struct {
			Host     string `json:"host"`
			Port     int    `json:"port"`
			User     string `json:"user"`
			Password string `json:"password"`
			Charset  string `json:"charset"`
		}{
			Host:     "172.22.0.2",
			Port:     3306,
			User:     "root",
			Password: "123456",
			Charset:  "utf8mb4",
		},
	}
	model.ConfigJson.UserIDList = []string{usrId}
	model.ConfigJson.EndDate = dataStr
	// 转换为JSON字符串
	jsonData, err := json.MarshalIndent(model.ConfigJson, "", "    ")
	if err != nil {
		fmt.Println("转换JSON数据失败:", err)
		return
	}
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录失败:", err)
		return
	}
	jsonFile := dir + "\\weibospider-demo\\conf\\config.json"
	// 写入JSON数据到文件
	err = os.WriteFile(jsonFile, jsonData, 0644)
}
