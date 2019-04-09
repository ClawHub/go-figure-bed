package utils

import (
	"github.com/Unknwon/com"
	"go-figure-bed/pkg/logging"
	"go.uber.org/zap"
	"os"
	"regexp"
	"strings"
	"time"
)

//文件类型
var picType = []string{"png", "jpg", "jpeg", "gif", "bmp"}

//通过 MimeType 信息获取文件名称
func GetFileNameByMimeType(info string) string {
	pat := `filename="(.*)"`
	res := regexp.MustCompile(pat)
	name := res.FindAllStringSubmatch(info, -1)
	return name[0][1]
}

//检查路径并且创建
func CheckPath(path string) {
	if _, err := os.Stat(path); err != nil {
		err = os.MkdirAll(path, 0775)
		if err != nil {
			logging.AppLogger.Error("Create Images store unsuccessful:", zap.Error(err))
			return
		}
	}
}

//验证文件后缀&文件MIME
func Validate(contentType string, fileName string) bool {
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

//获取虚拟URL
func GetVirtualUrl(name string) string {
	nowTime := com.DateT(time.Now(), "/YYYY/MM/DD/")
	suffix := GetRandomString(8) + "-" + name
	return nowTime + suffix
}
