package setting

//上传到本地
type Local struct {
	Open            bool
	StorageLocation string
	Link            string
}

//新浪
type Sina struct {
	//是否开启微博图床
	OpenSinaPicStore bool
	//用户名
	UserName string
	//密码
	PassWord string
	//新浪 Cookie 更新的频率,默认为3600s ,单位 s
	ResetSinaCookieTime int
	//新浪图床默认使用的尺寸大小 square,thumb150,orj360,orj480,mw690,mw1024,mw2048,small,bmiddle,large 、默认为large
	DefultPicSize string
}

//雅虎
type Flickr struct {
	//default size
	DefaultSize string
	//api_key
	Id               string
	ApiKey           string
	ApiSecret        string
	OauthToken       string
	OauthTokenSecret string
	//是否开启 flickr 图床 (此功能该可以在后台开启)
	OpenFlickrStore bool
}
