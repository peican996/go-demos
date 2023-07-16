package main

import (
	"demos/weibospider-demo/config/model"
	"demos/weibospider-demo/utils"
	"fmt"
)

func main() {
	testData()
}

func testData() {
	data := model.Config{
		UserIDList:          []string{},
		Filter:              1,
		SinceDate:           "2018-01-01",
		EndDate:             "2018-01-01 18:59",
		RandomWaitPages:     []int{1, 5},
		RandomWaitSeconds:   []int{6, 10},
		GlobalWait:          [][]int{{1000, 3600}, {500, 2000}},
		WriteMode:           []string{"txt", "mysql"},
		PicDownload:         0,
		VideoDownload:       0,
		FileDownloadTimeout: []int{5, 5, 10},
		ResultDirName:       0,
		Cookie:              "_T_WM=63b8f9b8c6511541c736b8c0a624a827; __e_inc=1; SCF=AjmpG2nEVUL3Uomo8ejN54uKCTTWgL_J4qkA4pH9OZ3SnxB8VcdphlXGD071pA5C9PlkbTpx-D52JvpLOsVIh-w.; SUB=_2A25JtgQsDeRhGeNO6VoQ9C7IyjmIHXVrWKxkrDV6PUJbktANLUmjkW1NTxjllnUwltMgZLuswXNR34OvWTpVlrDc; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9Wh57jVs4GfJR_d7uc2_yGeB5JpX5K-hUgL.Fo-7eonpSh5XeK-2dJLoIEBLxKqL1hBLB.qLxK.L1-2L12qLxK-LB.qL1heLxKnLBK2L12et; SSOLoginState=1689416828; ALF=1692008828; XSRF-TOKEN=875dd7; WEIBOCN_FROM=1110006030; mweibo_short_token=01b225c94f; MLOGIN=1; M_WEIBOCN_PARAMS=luicode%3D20000174%26uicode%3D20000174",
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
	result := utils.ComplementJudge(data)
	expected := true
	if result != expected {
		fmt.Printf("ComplementJudge(data) 返回结果错误，期望 %s，实际得到 %s", expected, result)
	}

}
