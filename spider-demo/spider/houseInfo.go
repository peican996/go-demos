package spider

import (
	"LianjiaSpider/redis"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/jinzhu/gorm"

	"LianjiaSpider/model"

	"LianjiaSpider/common"
	"github.com/gocolly/colly"
)

var (
	id            string
	totalPrice    string
	unitPrice     string
	roomInfo      string
	areaInfo      string
	areaName      string
	communityInfo string
	transaction   string
	textBody      string
)

func GetHouseUrl(db *gorm.DB, page int, districtName string) string {
	var houseInfo string
	c := colly.NewCollector(
		//colly.Async(true),并发
		colly.AllowURLRevisit(),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)
	c.SetRequestTimeout(time.Duration(120) * time.Second)
	err := c.Limit(&colly.LimitRule{DomainGlob: common.Properties.ErShouFang, Parallelism: 1})
	if err != nil {
		return ""
	} //Parallelism代表最大并发数
	c.OnRequest(func(r *colly.Request) {
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnHTML(".sellListContent>li", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			if strings.Contains(el.Attr("href"), "html") {
				GetHouseInfo(db, el.Attr("href"))
			}
		})
		houseInfo = e.ChildAttr("a", "href")
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
		err = c.Visit(common.Properties.ErShouFang + districtName + "/pg" + strconv.Itoa(page))
		if err != nil {
			return
		}
	})
	err = c.Visit(common.Properties.ErShouFang + districtName + "/pg" + strconv.Itoa(page))
	if err != nil {
		return ""
	}
	c.Wait()
	return houseInfo
}

func GetHouseInfo(db *gorm.DB, houseInfoUrl string) {
	var houseInfo model.HouseInfo
	var body model.Body
	c := colly.NewCollector(
		//colly.Async(true),并发
		colly.AllowURLRevisit(),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)
	c.SetRequestTimeout(time.Duration(120) * time.Second)
	err := c.Limit(&colly.LimitRule{DomainGlob: houseInfoUrl, Parallelism: 1})
	if err != nil {
		return
	} //Parallelism代表最大并发数
	c.OnRequest(func(r *colly.Request) {
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnHTML("body", func(e *colly.HTMLElement) {
		textBody = e.Text
		body.TextBody = textBody
	})
	c.OnHTML(".overview .content", func(e *colly.HTMLElement) {
		id = e.ChildText("div.aroundInfo > div.houseRecord > span.info")
		totalPrice = e.ChildText("div.price-container > div.price > span.total")
		unitPrice = e.ChildText("div.price-container > div.price > div.text > div.unitPrice > span.unitPriceValue")
		roomInfo = e.ChildText("div.houseInfo > div.room > div.mainInfo") + "," + e.ChildText("div.houseInfo > div.room > div.subInfo")
		areaInfo = e.ChildText("div.area > div.mainInfo") + "," + e.ChildText("div.area > div.subInfo")
		areaName = e.ChildText("div.aroundInfo > div.areaName > span.info")
		communityInfo = e.ChildText("div.aroundInfo > div.communityName > a.info") + "," + e.ChildText("div.aroundInfo > div.areaName > a")
		//var houseInfo := model.HouseInfo{}
		body.Id = getIdNum(id)
		houseInfo.Id = getIdNum(id)
		houseInfo.TotalPrice = totalPrice
		houseInfo.UnitPrice = encodedStr(unitPrice)
		houseInfo.RoomInfo = encodedStr(roomInfo)
		houseInfo.AreaInfo = encodedStr(areaInfo)
		houseInfo.AreaName = encodedStr(trimSpace(areaName))
		houseInfo.CommunityInfo = encodedStr(communityInfo)
	})
	c.OnHTML(".m-content .introContent .transaction", func(e *colly.HTMLElement) {
		transaction = e.ChildText("div.content > ul > li:nth-child(4) > span:nth-child(2)")
		houseInfo.Transaction = encodedStr(transaction)
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
		err = c.Visit(houseInfoUrl)
		if err != nil {
			return
		}
	})
	err = c.Visit(houseInfoUrl)
	if err != nil {
		return
	}
	c.Wait()
	redis.CacheValue(houseInfo)
	//err = db.Save(&houseInfo).Error
	//for err != nil {
	//	sellingInfo := model.HouseInfo{Id: getIdNum(id), TotalPrice: totalPrice, UnitPrice: unitPrice, RoomInfo: roomInfo, AreaInfo: areaInfo, AreaName: areaName, CommunityInfo: communityInfo, Transaction: transaction}
	//	err = db.Save(&sellingInfo).Error
	//}
	//writeToFile(getIdNum(id), textBody)
}

func getIdNum(id string) string {
	// 使用正则表达式匹配数字部分
	re := regexp.MustCompile(`\d+`)
	digits := re.FindAllString(id, -1)
	return digits[0]
}

// 获取非 NBSP 中文字符部分
func trimSpace(str string) string {
	result := ""
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) && r != '\u00A0' {
			result += string(r)
		}
	}
	re := regexp.MustCompile(`&nbsp;`)
	bb := re.ReplaceAllString(result, "")

	// 替换空白字符为空字符串
	re = regexp.MustCompile(`\s*`)
	cc := re.ReplaceAllString(bb, "")
	return cc
}

func writeToFile(fileName string, content string) {
	fileName = "D:\\lianjiehtml\\" + fileName + ".html"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("文件打开失败", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	_, err = f.WriteString(content)
	if err != nil {
		log.Println("文件写入失败", err)
		return
	}
}

func encodedStr(encodedStr string) string {
	//将字符串转换为utf-8
	encodeString, _ := strconv.Unquote(`"` + encodedStr + `"`)
	return encodeString
}
