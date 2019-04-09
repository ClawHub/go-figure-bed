package server

import (
	"encoding/json"
	"go-figure-bed/pkg/e/bed"
	"go-figure-bed/pkg/utils"
)

func UploadToAli(img []byte, imgInfo string, imgType string) string {
	url := "https://kfupload.alibaba.com/mupload"
	name := utils.GetFileNameByMimeType(imgInfo)

	file := &utils.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	//var header map[string]string
	data := utils.AliFormPost(file, url)
	ali := bed.AliResp{}
	_ = json.Unmarshal([]byte(data), &ali)
	return ali.Url
}
