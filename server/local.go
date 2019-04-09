package server

import (
	"bufio"
	"github.com/Unknwon/com"
	"go-figure-bed/pkg/logging"
	"go-figure-bed/pkg/setting"
	"go-figure-bed/pkg/utils"
	"go.uber.org/zap"
	"os"
	"time"
)

//图片上传到本地
func UpLoadToLocal(name string, fileContent []byte) (string, string, string, string) {
	//上传本地开关是否开启
	if !setting.BedSetting.Local.Open {
		return "", "", "", ""
	}
	//网站链接
	host := &setting.AppSetting.SiteUrl
	storeLocation := &setting.BedSetting.Local.StorageLocation
	softLink := &setting.BedSetting.Local.Link

	//修正URL
	utils.FormatUrl(softLink)
	utils.FormatUrl(host)
	utils.FormatUrl(storeLocation)

	//储存图片
	suffix := storeImage(*storeLocation, name, fileContent)
	url := *host + *softLink + suffix
	backup := *host + "backup/" + suffix
	randomStr := utils.GetRandomString(16)
	return url, backup, randomStr, *storeLocation + suffix
}

//储存图片
func storeImage(path string, n string, fileContent []byte) string {
	nowTime := com.DateT(time.Now(), "YYYY/MM/DD/")
	suffix := utils.GetRandomString(16) + "." + getImageSuffix(n)
	dir := path + nowTime
	file := dir + suffix
	//检查路径并且创建
	utils.CheckPath(dir)
	var f *os.File
	f, err := os.Create(file)
	if err != nil {
		logging.AppLogger.Error("File Create Error:", zap.Error(err))
	}
	w := bufio.NewWriter(f)
	_, err = w.Write(fileContent)
	if err != nil {
		logging.AppLogger.Error("File Create Error:", zap.Error(err))
	}
	_ = w.Flush()
	_ = f.Close()
	return nowTime + suffix
}

//获取图片后缀
func getImageSuffix(name string) (suffix string) {
	n := len(name)
	rs := []rune(name)
	suffix = string(rs[n-3 : n])
	if suffix == "peg" {
		suffix = "jpeg"
	}
	return suffix
}
