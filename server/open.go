package server

import (
	"encoding/json"
	"go-figure-bed/pkg/e/bed"
	"go-figure-bed/pkg/logging"
	"go-figure-bed/pkg/utils"
)

//存在 api 限制问题，暂时不考虑接入
func UpLoadToPublicSina(img []byte, imgName string, imgType string) string {
	url := "https://apis.yum6.cn/api/5bd44dc94bcfc?token=f07b711396f9a05bc7129c4507fb65c5"

	file := &utils.FormFile{
		Name:  imgName,
		Key:   "file",
		Value: img,
	}
	var header map[string]string
	data := utils.FormPost(file, url, header)
	open := bed.SinaPublicResponse{}
	_ = json.Unmarshal([]byte(data), &open)
	pid, ok := open.Data["pid"].(string)
	if !ok {
		logging.AppLogger.Error("上传公共图床出错")
		return ""
	}
	url = utils.CheckPid(pid, imgType)
	return url
}
