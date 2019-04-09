package server

import (
	"encoding/json"
	"go-figure-bed/pkg/e/bed"
	"go-figure-bed/pkg/utils"
)

func UploadToNetEasy(img []byte, imgInfo string, imgType string) string {
	url := "http://you.163.com/xhr/file/upload.json"
	name := utils.GetFileNameByMimeType(imgInfo)

	file := &utils.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := utils.FormPost(file, url, header)
	netEasy := bed.NetEasyResp{}

	_ = json.Unmarshal([]byte(data), &netEasy)
	return netEasy.Data[0]
}
