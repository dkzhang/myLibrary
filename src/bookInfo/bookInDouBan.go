package bookInfo

import (
	"fmt"
	"myUtils"
	"regexp"
)

type BookInDouBan struct {
	ID  string
	URL string
}

func (book *BookInDouBan) MakeURL() {
	book.URL = fmt.Sprintf(DOU_BAN_URL_FORMAT, book.ID)
}

func (book *BookInDouBan) ParseID() {
	re := regexp.MustCompile(`/([0-9]*)/?$`)
	match := re.FindStringSubmatch(book.URL)
	if match != nil {
		book.ID = myUtils.TrimBlank(match[1])
	}
}

const DOU_BAN_URL_FORMAT = `https://book.douban.com/subject/%s/`
