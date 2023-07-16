package utils

import (
	"demos/weibospider-demo/config/model"
	"fmt"
	"strings"
	"time"
)

func ComplementJudge(config model.Config) bool {
	if config.EndDate == "now" {
		return true
	}
	fmt.Println("config.SinceDate: ", config.SinceDate)
	if strings.Contains(getYearMonth(config.EndDate), config.SinceDate) {
		return true
	}
	return false
}

func getYearMonth(date string) string {
	layout := "2006-01-02 15:04"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("日期解析错误:", err)
		return ""
	}
	dateYearMonth := t.Format("2006-01-02")
	fmt.Println("日期:", dateYearMonth)
	return dateYearMonth
}
