package server

import (
	"bytes"
	"encoding/json"
	"go-figure-bed/pkg/e/bed"
	"go-figure-bed/pkg/logging"
	"go.uber.org/zap"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

//上传 SM 图床 返回图片 URL
func UploadToSmms(img []byte, imgName string) string {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	contentType := w.FormDataContentType()
	file, _ := w.CreateFormFile("smfile", imgName)
	_, _ = file.Write(img)
	_ = w.Close()
	req, _ := http.NewRequest("POST", "https://sm.ms/api/upload", body)
	req.Header.Set("Content-Type", contentType)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	sm := bed.SmResponse{}
	err := json.Unmarshal([]byte(string(data)), &sm)
	if err != nil {
		logging.AppLogger.Error("Upload To Smms fail", zap.Error(err))
		return ""
	}
	return string(sm.Data.Url)
}
