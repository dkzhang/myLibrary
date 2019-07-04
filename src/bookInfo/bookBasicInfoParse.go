package bookInfo

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strings"
)

type BookBasicInfoRegexp struct {
	FieldRegexp   *regexp.Regexp
	FieldFunction func(*BookBasicInfo) *string
	FiledName     string
}

var theBookBasicInfoRegexp = []BookBasicInfoRegexp{
	{FieldRegexp: regexp.MustCompile(`<span[^>]*>出版社:</span>([^<]*)<br/>`),
		FieldFunction: func(info *BookBasicInfo) *string {
			return &info.Publisher
		},
		FiledName: "出版社",
	},

	{FieldRegexp: regexp.MustCompile(`<span[^>]*>原作名:</span>([^<]*)<br/>`),
		FieldFunction: func(info *BookBasicInfo) *string {
			return &info.OriginalName
		},
		FiledName: "原作名",
	},

	{FieldRegexp: regexp.MustCompile(`<span[^>]*>出版年:</span>([0-9\-\s]*)<br/>`),
		FieldFunction: func(info *BookBasicInfo) *string {
			return &info.PublicationDate
		},
		FiledName: "出版年",
	},

	{FieldRegexp: regexp.MustCompile(`<span[^>]*>页数:</span>([0-9\s]*)<br/>`),
		FieldFunction: func(info *BookBasicInfo) *string {
			return &info.Pages
		},
		FiledName: "页数",
	},

	{FieldRegexp: regexp.MustCompile(`<span[^>]*>定价:</span>([0-9a-zA-Z\.\s]*)<br/>`),
		FieldFunction: func(info *BookBasicInfo) *string {
			return &info.Price
		},
		FiledName: "定价",
	},

	{FieldRegexp: regexp.MustCompile(`<span[^>]*>装帧:</span>([^<]*)<br/>`),
		FieldFunction: func(info *BookBasicInfo) *string {
			return &info.BindingAndLayout
		},
		FiledName: "装帧",
	},

	{FieldRegexp: regexp.MustCompile(`<span class="pl">丛书:</span>[^<]*<a[^>]*>([^<]*)</a><br/>`),
		FieldFunction: func(info *BookBasicInfo) *string {
			return &info.BookSeries
		},
		FiledName: "丛书",
	},

	{FieldRegexp: regexp.MustCompile(`<span[^>]*>ISBN:</span>([0-9\s]*)<br/>`),
		FieldFunction: func(info *BookBasicInfo) *string {
			return &info.ISBN
		},
		FiledName: "ISBN",
	},
}

func (bookBasicInfo *BookBasicInfo) ParseFromHtml(doc *goquery.Document) (err error) {
	//书名
	name := doc.Find("#wrapper > h1:nth-child(2) > span:nth-child(1)").Text()
	log.Println(name)
	bookBasicInfo.BookName = name

	//CSS选择器匹配字段
	bookBasicInfoSelection := doc.Find("div #info")

	bookBasicInfoSelection.Find("span").Each(func(i int, s1 *goquery.Selection) {
		key := s1.Find("span").Text()

		if strings.Trim(key, " ") == "作者" {
			s1.Find("a").Each(func(i int, s2 *goquery.Selection) {
				bookBasicInfo.Author = append(bookBasicInfo.Author, s2.Text())
			})
			log.Println(bookBasicInfo.Author)
		}

		if strings.Trim(key, " ") == "译者" {
			s1.Find("a").Each(func(i int, s2 *goquery.Selection) {
				bookBasicInfo.Translator = append(bookBasicInfo.Translator, s2.Text())
			})
			log.Println(bookBasicInfo.Translator)
		}
	})

	//正则表达式匹配字段
	fullHtml, err := bookBasicInfoSelection.Html()
	if err != nil {
		panic(err)
	}
	log.Println(fullHtml)

	err = bookBasicInfo.parseRegexp(fullHtml, theBookBasicInfoRegexp)
	if err != nil {
		return err
	}

	return nil
}

func (bookBasicInfo *BookBasicInfo) parseRegexp(fullHtml string, parseRegexp []BookBasicInfoRegexp) (err error) {
	for _, re := range parseRegexp {
		match := re.FieldRegexp.FindStringSubmatch(fullHtml)
		if match != nil {
			str := strings.Trim(match[1], " ")
			*(re.FieldFunction(bookBasicInfo)) = str
			log.Printf("%s：%s", re.FiledName, str)
		}
	}
	return nil
}
