package bookInfo

import (
	"github.com/PuerkitoBio/goquery"
	"myLibrary/src/myUtils"
	"regexp"
	"strconv"
	"strings"
)

func (douBanRating *DouBanRating) ParseFromHtml(doc *goquery.Document) (err error) {

	//CSS选择器匹配字段
	douBanRatingSelection := doc.Find("div #interest_sectl")

	fullHtml, err := douBanRatingSelection.Html()
	if err != nil {
		panic(err)
	}
	//log.Println(fullHtml)

	noRatingRegexp := regexp.MustCompile(`">评价人数不足</a>`)
	if len(noRatingRegexp.FindString(fullHtml)) == 0 {
		//评分可用
		douBanRating.HasRating = true

		//评分
		ratingStr := douBanRatingSelection.Find("strong").Text()
		douBanRating.Rating, _ = strconv.ParseFloat(strings.Trim(ratingStr, " "), 64)

		//评价人数
		ratingNumberText := douBanRatingSelection.Find("div.rating_sum").Text()
		re := regexp.MustCompile(`([0-9\s]*)人评价`)
		match := re.FindStringSubmatch(ratingNumberText)
		if match != nil {
			str := myUtils.TrimBlank(match[1])
			douBanRating.RatingNumber, _ = strconv.Atoi(str)
		}

		//各星人数
		douBanRatingSelection.Find("span.rating_per").Each(func(i int, s1 *goquery.Selection) {
			douBanRating.Stars[5-i], _ = strconv.ParseFloat(strings.Trim(s1.Text(), " %"), 64)
		})
	} else {
		//评分不可用
		douBanRating.HasRating = false
	}
	return nil
}
