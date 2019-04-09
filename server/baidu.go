package server

import (
	"bytes"
	"encoding/json"
	"go-figure-bed/pkg/e/bed"
	"go-figure-bed/pkg/logging"
	"go-figure-bed/pkg/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

//百度识图的接口
func UploadToBaidu(img []byte, imgInfo string) string {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	contentType := w.FormDataContentType()
	name := utils.GetFileNameByMimeType(imgInfo)
	file, _ := w.CreateFormFile("Filedata", name)
	_, _ = file.Write(img)
	_ = w.WriteField("file", "multipart")
	_ = w.Close()
	req, _ := http.NewRequest("POST", "https://api.uomg.com/api/image.baidu", body)
	req.Header.Set("Content-Type", contentType)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	baidu := bed.BaiduResp{}
	err := json.Unmarshal([]byte(string(data)), &baidu)
	if err != nil {
		logging.AppLogger.Error("Upload To Baidu fail", zap.Error(err))
		return ""
	}
	return string(baidu.ImgUrl)
}
