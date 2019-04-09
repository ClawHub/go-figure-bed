package server

import (
	"encoding/json"
	"go-figure-bed/pkg/e/bed"
	"go-figure-bed/pkg/utils"
)

func UploadToJueJin(img []byte, imgInfo string, imgType string) string {
	url := "https://cdn-ms.juejin.im/v1/upload?bucket=gold-user-assets"
	name := utils.GetFileNameByMimeType(imgInfo)

	file := &utils.FormFile{
		Name:  name,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := utils.FormPost(file, url, header)
	juejin := bed.JueJinResp{}
	_ = json.Unmarshal([]byte(data), &juejin)

	//神奇三断言 : )
	reJ, _ := juejin.D.(map[string]interface{})
	urls, _ := reJ["url"].(map[string]interface{})
	httpUrl, _ := urls["https"].(string)
	return httpUrl
}
