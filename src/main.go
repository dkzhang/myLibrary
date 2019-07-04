package main

import (
	"bookInfo"
	"strings"

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
	//res, err := client.Get("https://book.douban.com/subject/27015617/")
	res, err := client.Get("https://book.douban.com/subject/30424330/")
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

	GetBookBasicInfo(doc)

	/*
		// Find the review items
		doc.Find("div #info").Each(func(i int, s1 *goquery.Selection) {

			s1.Find("span").Each(func(i int, s2 *goquery.Selection) {
				key := s2.Find("span").Text()

				if s2.Text() == "出版社:" {
					value := s2.Text()
					fmt.Printf("Find key = %s, value = %s \n", key, value)
				}
				values := []string{}
				s2.Find("a").Each(func(i int, s3 *goquery.Selection) {
					values = append(values, s3.Text())
				})
				fmt.Printf("Find key = %s, values = %v \n", key, values)
			})
		})
	*/
}

func GetBookBasicInfo(doc *goquery.Document) bookInfo.BookBasicInfo {
	bbi := bookInfo.BookBasicInfo{}

	//书名
	name := doc.Find("#wrapper > h1:nth-child(2) > span:nth-child(1)").Text()
	log.Println(name)
	bbi.BookName = name

	bookBasicInfoSelection := doc.Find("div #info")

	bookBasicInfoSelection.Find("span").Each(func(i int, s1 *goquery.Selection) {
		key := s1.Find("span").Text()

		if strings.Trim(key, " ") == "作者" {
			s1.Find("a").Each(func(i int, s2 *goquery.Selection) {
				bbi.Author = append(bbi.Author, s2.Text())
			})
			log.Println(bbi.Author)
		}

		if strings.Trim(key, " ") == "译者" {
			s1.Find("a").Each(func(i int, s2 *goquery.Selection) {
				bbi.Translator = append(bbi.Translator, s2.Text())
			})
			log.Println(bbi.Translator)
		}
	})

	//正则表达式匹配字段
	fullHtml, err := bookBasicInfoSelection.Html()
	if err != nil {
		panic(err)
	}
	//log.Println(fullHtml)

	rePublicationDate := regexp.MustCompile(`<span[^>]*>出版年:</span>([0-9\-\s]*)<br/>`)
	matchPublication := rePublicationDate.FindStringSubmatch(fullHtml)
	if matchPublication != nil {
		bbi.PublicationDate = strings.Trim(matchPublication[1], " ")
		log.Printf("出版年：%s", bbi.PublicationDate)
	}

	rePages := regexp.MustCompile(`<span[^>]*>出版年:</span>([0-9\-\s]*)<br/>`)
	match := rePages.FindStringSubmatch(fullHtml)
	if match != nil {
		bbi.PublicationDate = strings.Trim(match[1], " ")
		log.Printf("出版年：%s", bbi.PublicationDate)
	}

	return bbi
}

var regexpPublicationDate = regexp.MustCompile(`<span[^>]*>出版年:</span>([0-9\-\s]*)<br/>`)

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
