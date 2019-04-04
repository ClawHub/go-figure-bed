package setting

import (
	"go-figure-bed/pkg/logging"
	"go.uber.org/zap"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	//最大上传的图片个数
	SiteUploadMaxNumber int
	//最大图片规格 MB
	SiteUploadMaxSize int64
	//Api 默认上传图床
	ApiDefault string
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

//上传到本地
type LocalStore struct {
	Open            bool
	StorageLocation string
	Link            string
}

var LocalStoreSetting = &LocalStore{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		logging.AppLogger.Fatal("setting.Setup, fail to parse 'conf/app.ini' ", zap.Error(err))
	}

	MapTo("app", AppSetting)
	MapTo("server", ServerSetting)
	MapTo("localStore", LocalStoreSetting)
	//图片规格 MB 2^20
	AppSetting.SiteUploadMaxSize = AppSetting.SiteUploadMaxSize << 20
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func MapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logging.AppLogger.Fatal("Cfg.MapTo Setting err' ", zap.Error(err))
	}
}
