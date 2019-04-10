package routers

import (
	"github.com/gin-gonic/gin"
	"go-figure-bed/pkg/setting"
	"go-figure-bed/routers/api"
)

//初始化路由
func InitRouter() *gin.Engine {

	//默认初始化 Gin
	r := gin.New()
	//Logger实例将日志写入gin.DefaultWriter的日志记录器中间件。
	r.Use(gin.Logger())

	//Recovery返回一个中间件，该中间件从任何恐慌中恢复，如果有500，则写入500。
	r.Use(gin.Recovery())
	//设置mode-----"debug","release","test"
	gin.SetMode(setting.ServerSetting.RunMode)

	//健康检查
	r.GET("/welcome", api.Welcome)

	//上传图片
	r.POST("/upload", api.UploadImage)
	//下载图片
	//r.StaticFS(setting.BedSetting.Link, http.Dir(setting.BedSetting.StorageLocation))

	//读取图片
	r.GET("/image/*mainUrl", api.GetImage)

	//检测并修复库中失效图片url
	r.GET("/autoCheckRepair", api.AutoCheckRepair)
	return r
}
