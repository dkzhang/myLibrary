package bookInfo

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
)

func (bookInfo *BookInfo) ParseFromHtmlDouBan(bookIDInDouBan string) (err error) {
	bookInfo.TheBookInDouBan.ID = bookIDInDouBan
	bookInfo.TheBookInDouBan.MakeURL()

	urli := url.URL{}
	urlproxy, _ := urli.Parse("http://proxy7.bj.petrochina:8080")
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}

	// Request the HTML page.
	res, err := client.Get(bookInfo.TheBookInDouBan.URL)
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

	/*
		err = bookInfo.TheBookBasicInfo.ParseFromHtml(doc)
		if err != nil {
			return err
		}

		err = bookInfo.TheDouBanRating.ParseFromHtml(doc)
		if err != nil {
			return err
		}

		err = bookInfo.TheBookCover.ParseFromHtml(doc)
		if err != nil {
			return err
		}

		imageFilePath := fmt.Sprintf("%s_%s.jpg",
			bookInfo.TheBookBasicInfo.BookName, bookInfo.TheBookBasicInfo.ISBN)
		err = bookInfo.TheBookCover.DownloadCover(imageFilePath)
		if err != nil {
			return err
		}
	*/

	bookInfo.TheBookIntroduce.ParseFromHtml(doc, bookInfo.TheBookInDouBan.ID)

	return nil
}
