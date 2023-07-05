package main

import (
	"LianjiaSpider/common"
	"LianjiaSpider/spider"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

var (
	district = [13]string{"tianhe", "yuexiu", "liwan", "haizhu", "panyu", "baiyun",
		"huangpugz", "conghua", "zengcheng", "huadou", "nansha", "nanhai", "shunde"}

	//district = [1]string{"tianhe"}
)

// https://gz.lianjia.com/ershoufang/ 广州市链家网爬虫
func main() {
	// 创建一个时间通道，每小时发送一次时间
	ticker := time.Tick(time.Hour)
	fmt.Println("start first spider task")
	workTask()
	fmt.Println("end spider task")
	// 循环监听时间通道
	for {
		fmt.Println("start spider task")
		select {
		case <-ticker:
			// 执行定时任务的操作
			fmt.Println("执行价格数据对比任务")
			fmt.Println("start spider task")
			workTask()
			fmt.Println("end spider task")
		}
	}
}

func workTask() {
	//初始化配置
	InitConfig()
	common.InitConfig()
	db := common.InitDB()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("close db error")
		}
	}(db)
	//var wgSelling sync.WaitGroup
	//var wgSold sync.WaitGroup
	var wgHouseInfo sync.WaitGroup
	for _, districtName := range district {
		totalSellingPage := spider.GetSellingPageSpider(districtName)
		for page := 1; page < totalSellingPage; page++ {
			wgHouseInfo.Add(1)
			time.Sleep(time.Duration(100) * time.Millisecond)
			go func(page int) {
				defer wgHouseInfo.Done()
				spider.GetHouseUrl(db, page, districtName)
			}(page)
		}
	}
	wgHouseInfo.Wait()
}

// InitConfig 初始化配置函数
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
