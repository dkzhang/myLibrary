package bookInfo

import (
	"github.com/PuerkitoBio/goquery"
	"log"
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
		log.Println(ratingNumberText)
		re := regexp.MustCompile(`([0-9\s]*)人评价`)
		match := re.FindStringSubmatch(ratingNumberText)
		if match != nil {
			str := strings.Trim(match[1], " \f\n\r\t\v")
			douBanRating.RatingNumber, _ = strconv.Atoi(str)
		}

		//各星人数
		douBanRatingSelection.Find("span.rating_per").Each(func(i int, s1 *goquery.Selection) {
			log.Println(i, ":", s1.Text())
			douBanRating.Stars[5-i], _ = strconv.ParseFloat(strings.Trim(s1.Text(), " %"), 64)
		})
	} else {
		//评分不可用
		douBanRating.HasRating = false
	}
	/*
			douBanRatingSelection.Find("span").Each(func(i int, s1 *goquery.Selection) {
			key := s1.Find("span").Text()

			if strings.Trim(key, " ") == "作者" {
				s1.Find("a").Each(func(i int, s2 *goquery.Selection) {
					bookBasicInfo.Author = append(bookBasicInfo.Author, s2.Text())
				})
				log.Println(bookBasicInfo.Author)
			}
		}
	*/

	return nil
}
