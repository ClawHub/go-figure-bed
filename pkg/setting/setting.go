package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	//最大上传的图片个数
	SiteUploadMaxNumber int
	//最大图片规格 MB
	SiteUploadMaxSize int64
	//根节点图床
	RootSiteApi []string
	//其他节点
	OtherSiteApi []string
	//网站链接
	SiteUrl string
}

var AppSetting = &App{}

//服务相关
type Server struct {
	ProjectName  string
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

//图床配置
type Bed struct {
	Local
	Sina
	Flickr
}

var BedSetting = &Bed{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("app.ini")
	if err != nil {
		log.Fatal("setting.Setup, fail to parse 'app.ini' ", err)
	}

	MapTo("app", AppSetting)
	MapTo("server", ServerSetting)
	//不指定Section
	MapToRoot(BedSetting)
	//图片规格 MB 2^20
	AppSetting.SiteUploadMaxSize = AppSetting.SiteUploadMaxSize << 20
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func MapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatal("Cfg.MapTo Setting err' ", err)
	}
}
func MapToRoot(v interface{}) {
	err := cfg.MapTo(v)
	if err != nil {
		log.Fatal("Cfg.MapToRoot Setting err' ", err)
	}
}
