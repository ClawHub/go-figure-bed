package api

import (
	"fmt"
	"go-figure-bed/pkg/app"
	"go-figure-bed/pkg/e"
	"go-figure-bed/pkg/logging"
	"go-figure-bed/pkg/setting"
	"go-figure-bed/pkg/utils"
	"go-figure-bed/server"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//增加标签
type ImageForm struct {
	Apis []string `form:"apis" valid:"Required"`
}

func UploadImage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ImageForm
	)
	//绑定以及校验
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
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
	legal := validate(h.Header.Get("Content-Type"), h.Filename)
	if !legal {
		logging.HTTPLogger.Error("file type err")
		appG.Response(http.StatusInternalServerError, e.ERROR_FILE_TYPE, nil)
		return

	}
	c.Set("email", "ssss")
	//获取email
	email := c.GetString("email")
	if email == "" {
		logging.HTTPLogger.Error("get email fail")
		appG.Response(http.StatusInternalServerError, e.ERROR_FILE_UPLOAD, nil)
		return
	}
	//处理文件上传
	urls := HandleApis(email, form.Apis, f, h)
	//如果没有返回值
	//if !strings.HasPrefix(url, "http") {
	//	logging.HTTPLogger.Error("can not get img url")
	//	appG.Response(http.StatusInternalServerError, e.ERROR_CAN_NOT_GET_IMG_URL, nil)
	//	return
	//}
	//返回图片地址与文件名称
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"urls":     urls,
		"fileName": h.Filename,
	})
}

//处理文件上传
func HandleApis(email string, apis []string, f multipart.File, h *multipart.FileHeader) (urls []string) {
	imgMime := h.Header.Get("Content-Type")
	imgInfo := h.Header.Get("Content-Disposition")
	//读取文件名称
	imgName := utils.GetFileNameByMimeType(imgInfo)
	size := h.Size
	fileContent := make([]byte, size)
	_, _ = f.Read(fileContent)
	var url string
	//如果输入的apis内容为空
	if apis[0] == "" {
		//API默认图床选择
		switch setting.AppSetting.ApiDefault {
		case "Local":
			showUrl, backUrl, del, path := server.UpLoadToLocal(imgName, fileContent)
			url = showUrl
			fmt.Println(showUrl, backUrl, del, path)
		case "SouGou":
			url = server.UpLoadToSouGou(fileContent)
			proxyUrl := setting.AppSetting.SiteUrl + "api/proxy?url=" + url
			fmt.Println(proxyUrl)
		case "Sina":
			if setting.BedSetting.Sina.OpenSinaPicStore == false {
				url = ""
			} else {
				url = server.UpLoadToSina(fileContent, imgMime)
			}
			fmt.Println(url)
		case "Smms":
			url = server.UploadToSmms(fileContent, imgInfo)
			fmt.Println(url)
		case "CC":
			durl, del := server.UploadToCC(fileContent, imgInfo, imgMime)
			url = durl
			fmt.Println(url, del)
		case "Flickr":
		case "Baidu":
		case "Qihoo":
		case "NetEasy":
		case "Jd":
		case "Ali":
		case "Open":
		}
	}
	//迭代apis
	for _, api := range apis {
		//选中api
		switch api {
		case "Local":
		case "SouGou":
		case "Sina":
		case "Smms":
		case "CC":
		case "Flickr":
		case "Baidu":
		case "Qihoo":
		case "NetEasy":
		case "Jd":
		case "Ali":
		case "Open":
		default:

		}
		//放入结果中
		urls = append(urls, url)
	}
	return
}

//文件类型
var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}

//验证文件后缀&文件MIME
func validate(contentType string, fileName string) bool {
	//首先检测文件的后缀
	illegality := false
	for _, pType := range picType {
		if strings.HasSuffix(fileName, pType) {
			illegality = true
			break
		}
	}
	//然后检测 MIME 类型
	if strings.HasPrefix(contentType, "image") && illegality {
		for _, pType := range picType {
			if strings.HasSuffix(contentType, pType) {
				return true
			}
		}

	}
	return false
}
