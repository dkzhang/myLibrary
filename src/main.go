package main

import (
	"log"
	"myLibrary/src/bookInfo"
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
}
