package setting

//上传到本地
type LocalStore struct {
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

//Sina 图床 json
type SinaMsg struct {
	Code string   `json:"code"`
	Data SinaData `json:"data"`
}

type SinaData struct {
	Count int      `json:"count"`
	Data  string   `json:"data"`
	Pics  SinaPics `json:"pics"`
}

type SinaPics struct {
	Pic_1 picInfo `json:"pic_1"`
}

type picInfo struct {
	Width  int    `json:"width"`
	Size   int    `json:"size"`
	Ret    int    `json:"ret"`
	Height int    `json:"height"`
	Name   string `json:"name"`
	Pid    string `json:"pid"`
}

type SinaError struct {
	Retcode string `json:"retcode"`
	Reason  string `json:"reason"`
}
