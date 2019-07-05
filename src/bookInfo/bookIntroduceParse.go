package bookInfo

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func (bookIntro *BookIntroduce) ParseFromHtml(doc *goquery.Document, bookIDInDouBan string) (err error) {

	//先检查是否是带“展开全部”的内容简介
	contentIntro := doc.Find("#link-report > span.all.hidden > div > div.intro").Text()
	if len(contentIntro) != 0 {
		bookIntro.ContentIntroduce = contentIntro
	} else {
		//不带“展开全部”的内容简介
		bookIntro.ContentIntroduce = doc.Find("#link-report > div > div.intro").Text()
	}

	//先检查是否是带“展开全部”的作者简介
	authorIntro := doc.Find("div.indent > span.all.hidden > div.intro").Text()
	if len(authorIntro) != 0 {
		bookIntro.AuthorIntroduce = authorIntro
	} else {
		//不带“展开全部”的作者简介
		bookIntro.AuthorIntroduce = doc.Find("div.indent:nth-child(4) > div > div.intro").Text()
	}

	//完整版的目录
	bookIntro.Contents = doc.Find(fmt.Sprintf("#dir_%s_full", bookIDInDouBan)).Text()

	return nil
}
