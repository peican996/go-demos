package _type

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func GetVideo(urlVideo string) {
	url := urlVideo // 替换成你要爬取的网站URL

	c := colly.NewCollector()

	c.OnHTML("video", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		if strings.HasPrefix(src, "http") {
			fmt.Println("Video URL:", src)
		} else {
			fmt.Println("Video URL:", url+src)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
