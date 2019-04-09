package api

import (
	"github.com/gin-gonic/gin"
	"go-figure-bed/pkg/app"
	"go-figure-bed/pkg/e"
	"go-figure-bed/pkg/logging"
	"go-figure-bed/pkg/setting"
	"go-figure-bed/pkg/utils"
	"go-figure-bed/service"
	"go.uber.org/zap"
	"net/http"
)

func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	f, h, err := c.Request.FormFile("image")
	//判断上传的文件是否为空
	if f == nil {
		logging.HTTPLogger.Error("file is empty", zap.Error(err))
		appG.Response(http.StatusInternalServerError, e.ERROR_FILE_IS_EMPTY, nil)
		return
	}
	//判断文件是否太大
	if h.Size > setting.AppSetting.SiteUploadMaxSize {
		logging.HTTPLogger.Error("error file is too large", zap.Error(err))
		appG.Response(http.StatusInternalServerError, e.ERROR_FILE_IS_TOO_LARGE, nil)
		return
	}
	defer f.Close()
	//判断是否有错误
	if err != nil {
		logging.HTTPLogger.Error("File Upload Err", zap.Error(err))
		appG.Response(http.StatusInternalServerError, e.ERROR_FILE_UPLOAD, nil)
		return
	}
	//验证文件类型
	legal := utils.Validate(h.Header.Get("Content-Type"), h.Filename)
	if !legal {
		logging.HTTPLogger.Error("file type err")
		appG.Response(http.StatusInternalServerError, e.ERROR_FILE_TYPE, nil)
		return

	}
	//获取email
	email := c.GetString("email")
	if email == "" {
		//游客登陆
		email = "guest"
	}
	//处理文件上传
	url := service.Handle(email, f, h)
	//如果没有返回值
	if url == "" {
		logging.HTTPLogger.Error("can not get img url")
		appG.Response(http.StatusInternalServerError, e.ERROR_CAN_NOT_GET_IMG_URL, nil)
		return
	}
	//返回图片地址与文件名称
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"url":      url,
		"fileName": h.Filename,
	})
}
