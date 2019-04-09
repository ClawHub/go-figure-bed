package server

import (
	"encoding/json"
	"go-figure-bed/pkg/e/bed"
	"go-figure-bed/pkg/utils"
)

func UploadToNetEasy(img []byte, imgName string, imgType string) string {
	url := "http://you.163.com/xhr/file/upload.json"

	file := &utils.FormFile{
		Name:  imgName,
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
