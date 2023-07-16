package main

import (
	"bufio"
	"demos/weibospider-demo/common"
	"demos/weibospider-demo/config/model"
	"demos/weibospider-demo/service"
	"demos/weibospider-demo/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

var userId string

var retryTimes int

var preTime string

var currentTime string

var file string

var dataString string

func init() {
	common.InitConfig()
}

func main() {
	fmt.Println("开始爬取微博数据...")
	for _, id := range common.Info.UserIds {
		retryTimes = 0
		if len(id) == 10 {
			userId = id
			file = utils.FindFile(userId)
			dataString = utils.ObtainData(file)
			execSpider()
			uniqueLines(file)
		}
	}
}

func execSpider() {
	for {
		retryTimes++
		jsonFile := common.WorkPath + "\\weibospider-demo\\conf\\config.json"
		utils.RenewConfigJson(userId, dataString)
		preTime = model.ConfigJson.EndDate
		if judgeComplement(model.ConfigJson.EndDate) {
			fmt.Println("end data is greater than since data in the same week")
			break
		}
		// 当end_time和since_date相等时，跳出循环
		if strings.Contains(preTime, model.ConfigJson.SinceDate) {
			fmt.Println("end data is equal to since data")
			break
		}
		service.WeiboService(jsonFile)
		utils.RenewConfigJson(userId, dataString)
		currentTime = model.ConfigJson.EndDate
		time.Sleep(15 * time.Second)
		if preTime == currentTime {
			retryTimes++
			fmt.Println("retryTimes: ", retryTimes)
		} else if strings.Contains(preTime, model.ConfigJson.SinceDate) {
			break
		}
		fmt.Println("重试中...")
		if retryTimes >= 10 {
			fmt.Println("end data: " + model.ConfigJson.EndDate)
			break
		}
	}
}

func judgeComplement(data string) bool {
	// 假设有两个日期的字符串入参
	comparisonValue := "2018-01-01"
	fmt.Println("data: ", data)
	if strings.Contains(data, "now") {
		data = time.Now().Format("2006-01-02 15:04")
	}

	// 使用 "2006-01-02 15:04" 格式解析日期时间字符串为 time.Time 类型
	t, _ := time.Parse("2006-01-02 15:04", data)

	// 将时间格式化为 "2006-01-02" 形式
	formattedDate := t.Format("2006-01-02")

	// 使用 "2006-01-02" 格式解析日期字符串为 time.Time 类型
	t1, _ := time.Parse("2006-01-02", comparisonValue)
	t2, _ := time.Parse("2006-01-02", formattedDate)

	// 计算 t1 和 t2 之间的持续时间
	duration := t2.Sub(t1)

	// 创建一个表示七天的持续时间
	sevenDays := 7 * 24 * time.Hour

	// 检查持续时间是否小于等于七天
	if duration <= sevenDays {
		fmt.Println("t1 和 t2 相差七天之内")
		return true
	} else {
		fmt.Println("t1 和 t2 相差超过七天")
		return false
	}
}

// uniqueLines  file data
func uniqueLines(filePath string) {
	uniqueLineValues := make([]string, 0)
	lineMap := make(map[string]bool)
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件失败:", err)
	}
	defer file.Close()
	// 使用bufio.Scanner读取文件内容
	scanner := bufio.NewScanner(file)
	tempLine := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			line += "\n"
			tempLine += line
		} else if !lineMap[tempLine] {
			lineMap[tempLine] = true
			uniqueLineValues = append(uniqueLineValues, tempLine)
			tempLine = ""
		} else {
			fmt.Println("重复行：", tempLine)
			tempLine = ""
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	output := strings.Join(uniqueLineValues, "\n") + "\n"
	err = os.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Data written to test.txt successfully.")
}
