package server

import (
	"encoding/json"
	"go-figure-bed/pkg/e/bed"
	"go-figure-bed/pkg/logging"
	"go-figure-bed/pkg/utils"
	"go.uber.org/zap"
)

func UploadToCC(img []byte, imgName string, imgType string) (string, string) {
	url := "https://upload.cc/image_upload"

	file := &utils.FormFile{
		Name:  imgName,
		Key:   "uploaded_file[]",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := utils.FormPost(file, url, header)
	cc := bed.CCResp{}
	err := json.Unmarshal([]byte(data), &cc)
	if err != nil {
		logging.AppLogger.Error("UploadToCC fail ", zap.Error(err))
		return "", ""
	}
	mj, _ := cc.SuccessImage[0].(map[string]interface{})
	smj, _ := mj["url"].(string)
	del, _ := mj["delete"].(string)

	url = "https://upload.cc/" + smj

	deleteJson := `[{"path":"` + smj + `",key":"` + del + `"}]`
	return url, deleteJson
}
