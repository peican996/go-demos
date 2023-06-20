package main

import (
	"log"

	"github.com/gocolly/colly"
)

//func main() {
//	ScripDouban2()
//}

func ScripDouban2() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0"),
	)

	err := c.Limit(&colly.LimitRule{DomainGlob: "*.douban.*", Parallelism: 5})
	if err != nil {
		return
	}

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML(".item", func(e *colly.HTMLElement) {
		//log.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
		//	strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))
		log.Println(e.ChildAttr("a", "href"))
		//log.Println(e.ChildAttr(".info .hd .a", "href"))
		//log.Println(e.DOM.Find("img.src").Eq(0).Text())
		//log.Println(e.DOM.Find("img.alt").Eq(0).Text())
	})

	//c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
	//	e.Request.Visit(e.Attr("href"))
	//})

	err = c.Visit("https://movie.douban.com/top250?start=0&filter=")
	if err != nil {
		return
	}
	c.Wait()
}
