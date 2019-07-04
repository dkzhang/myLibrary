package main

import (
	"bookInfo"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

func ExampleScrape() {

	urli := url.URL{}
	urlproxy, _ := urli.Parse("http://proxy7.bj.petrochina:8080")
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}

	// Request the HTML page.
	res, err := client.Get("https://book.douban.com/subject/27015617/")
	//res, err := client.Get("https://book.douban.com/subject/30281411/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	theBookInfo := bookInfo.BookInfo{}
	theBookInfo.TheBookBasicInfo.ParseFromHtml(doc)
	theBookInfo.TheDouBanRating.ParseFromHtml(doc)

	log.Printf("%v", theBookInfo.TheBookBasicInfo)
	log.Printf("%v", theBookInfo.TheDouBanRating)
}

func main() {
	ExampleScrape()
	//TestRegexp()
}
func TestRegexp() {
	//	const text = `My email is ccmoust@gmail.com@abc.com
	//email1 is abc@def.org
	//email is kkk@qq.com
	//`
	//
	//	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)

	text := ` 80.00元`
	re := regexp.MustCompile(`[0-9]+`)
	match := re.FindAllString(text, -1)
	fmt.Println(match)
}
