package bookInfo

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"myLibrary/src/downloadFile"
)

func (bookCover *BookCover) ParseFromHtml(doc *goquery.Document) (err error) {
	//CSS选择器匹配字段
	bookCoverSelection := doc.Find("a.nbg")

	var exists bool
	bookCover.Url, exists = bookCoverSelection.Attr("href")

	if exists == false {
		return fmt.Errorf("can not find the book cover image url")
	} else {
		return nil
	}
}

func GenerateImageFilePath(bookInfo *BookInfo) string {
	return fmt.Sprintf("/BookCover/%s_%s.jpg",
		bookInfo.TheBookInDouBan.ID, bookInfo.TheBookBasicInfo.ISBN)
}

func (bookCover *BookCover) DownloadCover(imageFilePath string) (err error) {
	err = downloadFile.HttpDownloadFile(bookCover.Url, imageFilePath)
	if err == nil {
		bookCover.ImageFilePath = imageFilePath
	}
	return err
}
