package redis

import (
	"LianjiaSpider/model"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

func init() {
}

func CacheValue(houseInfo model.HouseInfo) {
	var preHouseInfo = model.HouseInfo{}
	// 序列化为JSON字符串
	houseInfoJSON, err := json.Marshal(houseInfo)
	if err != nil {
		log.Fatal("序列化用户数据失败:", err)
	}
	str := Client.Get(fmt.Sprintf("houserID:%s", houseInfo.Id))
	if str.Val() != "" {
		// 解析JSON数据到结构体
		err := json.Unmarshal([]byte(str.Val()), &preHouseInfo)
		if err != nil {
			fmt.Println("解析JSON失败:", err)
		}
		numUnitPrice, _ := strconv.Atoi(getIdNum(houseInfo.UnitPrice))
		numTmpUnitPrice, _ := strconv.Atoi(getIdNum(preHouseInfo.UnitPrice))
		priceChangeValue := numUnitPrice - numTmpUnitPrice
		if priceChangeValue != 0 {
			log.Printf("%s价格变动,原价为%s,现价格为%s,变动值为%s", houseInfo.Id, preHouseInfo.UnitPrice, houseInfo.UnitPrice, strconv.Itoa(priceChangeValue))
		}
	} else {
		// 设置到Redis缓存中
		err = Client.Set(fmt.Sprintf("houserID:%s", houseInfo.Id), string(houseInfoJSON), time.Hour).Err()
		if err != nil {
			log.Fatal("设置缓存失败:", err)
		}
	}
}

// 获取数字部分
func getIdNum(id string) string {
	// 使用正则表达式匹配数字部分
	re := regexp.MustCompile(`\d+`)
	digits := re.FindAllString(id, -1)
	return digits[0]
}
