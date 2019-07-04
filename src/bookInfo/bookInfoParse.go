package bookInfo

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func (bookInfo *BookInfo) ParseFromHtml(doc *goquery.Document) (err error) {
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

	return nil
}
