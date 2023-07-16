package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ObtainData(filePath string) string {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return "now"
	}
	defer file.Close()
	// 使用bufio.Scanner读取文件内容
	scanner := bufio.NewScanner(file)
	lastPublishTime := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "发布时间") {
			lastPublishTime = extractTime(line)
		}
	}
	// 检查是否发生错误
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件失败:", err)
		return "now"
	}
	// 输出最后一次发布时间
	fmt.Println("最后一次发布时间:", lastPublishTime)
	if lastPublishTime == "" {
		return "now"
	}
	return lastPublishTime
}

// 提取时间字符串的函数
func extractTime(line string) string {
	// 假设时间格式为：最后一次发布时间：2023-03-16 10:38
	parts := strings.Split(line, "：")
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1])
	}
	return "now"
}
