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
}

func (bookBasicInfo *BookBasicInfo) parse(doc *goquery.Document) (BookBasicInfo, error) {
	//书名
	name := doc.Find("#wrapper > h1:nth-child(2) > span:nth-child(1)").Text()
	log.Println(name)
	bookBasicInfo.BookName = name

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

	err = bookBasicInfo.parseRegexp(fullHtml)
	if err != nil {
		return *bookBasicInfo, err
	}

	return *bookBasicInfo, nil
}

/*
type BookBasicInfoRegexp struct {
	RegexpPublicationDate  *regexp.Regexp
	RegexpPages            *regexp.Regexp
	RegexpPrice            *regexp.Regexp
	RegexpBindingAndLayout *regexp.Regexp
}

func (parseRegexp *BookBasicInfoRegexp) init(){
	parseRegexp.RegexpPublicationDate = regexp.MustCompile(`<span[^>]*>出版年:</span>([0-9\-\s]*)<br/>`)
	parseRegexp.RegexpPages = regexp.MustCompile(`<span[^>]*>页数:</span>([0-9\s]*)<br/>`)
	parseRegexp.RegexpPrice = regexp.MustCompile(`<span[^>]*>定价:</span>([0-9a-zA-Z\.\s]*)<br/>`)
	parseRegexp.RegexpBindingAndLayout = regexp.MustCompile(`<span[^>]*>装帧:</span>([^<]*)<br/>`)
}

var once sync.Once
var theBookBasicInfoRegexp *BookBasicInfoRegexp = nil
func GetBookBasicInfoRegexp() *BookBasicInfoRegexp{
	once.Do(func() {
		theBookBasicInfoRegexp = &BookBasicInfoRegexp{}
		theBookBasicInfoRegexp.init()
	})
	return theBookBasicInfoRegexp
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (bookBasicInfo *BookBasicInfo) parse(doc *goquery.Document) (BookBasicInfo, error){
	//书名
	name := doc.Find("#wrapper > h1:nth-child(2) > span:nth-child(1)").Text()
	log.Println(name)
	bookBasicInfo.BookName = name

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
	if err !=nil{
		panic(err)
	}

	err = bookBasicInfo.parseRegexp(fullHtml)
	if err!=nil{
		return *bookBasicInfo, err
	}

	return *bookBasicInfo,nil
}

func (bookBasicInfo *BookBasicInfo) parseRegexp(fullHtml string) (err error){
	parseRegexp := GetBookBasicInfoRegexp()

	matchPublication := parseRegexp.RegexpPublicationDate.FindStringSubmatch(fullHtml)
	if matchPublication != nil{
		bookBasicInfo.PublicationDate = strings.Trim(matchPublication[1], " ")
		log.Printf("出版年：%s", bookBasicInfo.PublicationDate)
	}
	return nil
}

func regexpMatch(fullHtml string, re *regexp.Regexp, keyName string) (str string){
	match := re.FindStringSubmatch(fullHtml)
	if match != nil{
		str = strings.Trim(match[1], " ")
		log.Printf("%s：%s", keyName, str)
	}
	return
}


func UseBookBasicInfo(doc *goquery.Document) (BookBasicInfo, error){
	bbip := BookBasicInfo{doc: doc}
	return bbip.parse()
}

*/
