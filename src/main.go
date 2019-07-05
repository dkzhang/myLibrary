package main

import (
	"bookInfo"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
)

func ExampleScrape() {

	theBookInfo := bookInfo.BookInfo{}
	//theBookInfo.parseFromHtmlDouBan("27015617")
	//theBookInfo.parseFromHtmlDouBan("27665114")
	theBookInfo.ParseFromHtmlDouBanID("27133480", true)

	log.Printf("%v", theBookInfo.TheBookBasicInfo)
	log.Printf("%v", theBookInfo.TheDouBanRating)

	log.Printf("ContentIntroduce: %s", theBookInfo.TheBookIntroduce.ContentIntroduce)
	log.Printf("AuthorIntroduce: %s", theBookInfo.TheBookIntroduce.AuthorIntroduce)
	log.Printf("Tags: %s", theBookInfo.TheBookTagAndRec.Tags)
}

func main() {
	ExampleScrape()
	//TestSearch()
}
func TestSearch() {
	urli := url.URL{}
	urlproxy, _ := urli.Parse("http://proxy7.bj.petrochina:8080")
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}

	// Request the HTML page.
	res, err := client.Get("https://book.douban.com/subject_search?search_text=9787115460158&cat=1001")
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

	log.Println(doc.Html())
	log.Println(doc.Find(".title-text").Attr("href"))
}
