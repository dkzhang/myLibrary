package bookInfo

type BookBasicInfo struct {
	BookName         string   //书名
	Author           []string //作者
	Publisher        string   //出版社
	Producer         []string //出品方
	OriginalName     string   //原作名
	Translator       []string //译者
	PublicationDate  string   //出版年
	Pages            string   //页数
	Price            string   //定价
	BindingAndLayout string   //装帧
	BookSeries       string   //丛书
	ISBN             string   //ISBN
}
