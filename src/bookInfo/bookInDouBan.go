package bookInfo

import "fmt"

type BookInDouBan struct {
	ID  string
	URL string
}

func (book *BookInDouBan) MakeURL() {
	book.URL = fmt.Sprintf(DOU_BAN_URL_FORMAT, book.ID)
}

const DOU_BAN_URL_FORMAT = `https://book.douban.com/subject/%s/`
