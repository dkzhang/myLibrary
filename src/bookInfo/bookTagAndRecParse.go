package bookInfo

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"myUtils"
)

func (bookTagAndRec *BookTagAndRec) ParseFromHtml(doc *goquery.Document) (err error) {

	//CSS选择器匹配标签字段
	tagSelection := doc.Find("#db-tags-section > div.indent")

	tagSelection.Find("span").Each(func(i int, s *goquery.Selection) {
		bookTagAndRec.Tags = append(bookTagAndRec.Tags, myUtils.TrimBlank(s.Text()))
	})

	//CSS选择器匹配推荐书籍字段
	recSelection := doc.Find("#db-rec-section > div.content.clearfix")

	recSelection.Find("dl > dt > a").Each(func(i int, s *goquery.Selection) {
		log.Println(s.Attr("href"))
		url, exists := s.Attr("href")

		if exists == true && len(url) != 0 {
			book := BookInfo{}
			book.ParseFromHtmlDouBanURL(url, false)

			bookTagAndRec.RecommendBooks = append(bookTagAndRec.RecommendBooks, book)
		}
	})
	return nil
}
