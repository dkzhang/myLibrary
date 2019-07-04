package bookInfo

type DouBanRating struct {
	HasRating    bool    //是否有评分（按豆瓣规则，评价人数不足时，没有评分）
	Rating       float64 //豆瓣评分
	RatingNumber int     //评分人数
	Stars        [6]float64
}
