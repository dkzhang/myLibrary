package bookInfo

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func (bookInfo *BookInfo) ParseFromHtmlDouBanID(bookIDInDouBan string, isDeepParse bool) (err error) {
	bookInfo.TheBookInDouBan.ID = bookIDInDouBan
	bookInfo.TheBookInDouBan.MakeURL()
	return bookInfo.parseFromHtmlDouBan(isDeepParse)
}

func (bookInfo *BookInfo) ParseFromHtmlDouBanURL(bookURLInDouBan string, isDeepParse bool) (err error) {
	bookInfo.TheBookInDouBan.URL = bookURLInDouBan
	bookInfo.TheBookInDouBan.ParseID()
	return bookInfo.parseFromHtmlDouBan(isDeepParse)
}

func (bookInfo *BookInfo) parseFromHtmlDouBan(isDeepParse bool) (err error) {
	/*
		urli := url.URL{}
		urlproxy, _ := urli.Parse("http://proxy7.bj.petrochina:8080")
		client := http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(urlproxy),
			},
		}
	*/
	client := http.Client{}

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

	err = bookInfo.TheBookCover.DownloadCover(GenerateImageFilePath(bookInfo))
	if err != nil {
		return err
	}

	bookInfo.TheBookIntroduce.ParseFromHtml(doc, bookInfo.TheBookInDouBan.ID)

	if isDeepParse == true {
		bookInfo.TheBookTagAndRec.ParseFromHtml(doc)
	}

	return nil
}
