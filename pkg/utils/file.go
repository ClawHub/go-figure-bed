package utils

import (
	"go-figure-bed/pkg/logging"
	"go.uber.org/zap"
	"os"
	"regexp"
)

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
