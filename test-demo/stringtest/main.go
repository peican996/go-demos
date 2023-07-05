package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func main() {
	str := "108404578436举报"

	// 使用正则表达式匹配数字部分
	re := regexp.MustCompile(`\d+`)
	digits := re.FindAllString(str, -1)

	// 输出截取到的数字部分
	for _, digit := range digits {
		fmt.Println(digit)
	}

	str = "天河 天河公园"
	//fmt.Println(trimSpace(str))

	//str = "天河 天河公园"

	// 获取非 NBSP 中文字符部分
	result := ""
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) && r != '\u00A0' {
			result += string(r)
		}
	}

	fmt.Println(result) // 输出 "天河天河公园"
}

// 去除 和空格
func trimSpace(str string) string {
	// 使用正则表达式匹配数字部分
	re := regexp.MustCompile(`\s+`)
	digits := re.FindAllString(str, -1)
	return digits[0]
}
