package server

import (
	"go-figure-bed/pkg/utils"
	"regexp"
)

func UploadToQihoo(img []byte, imgInfo string, imgType string) string {
	url := "http://st.so.com/stu"
	name := utils.GetFileNameByMimeType(imgInfo)

	file := &utils.FormFile{
		Name:  name,
		Key:   "upload",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := utils.FormPost(file, url, header)
	var re = regexp.MustCompile(`(?m)data-imgkey="(.*)"`)
	imgKey := re.FindAllStringSubmatch(data, -1)[0][1]
	url = "https://ps.ssl.qhmsg.com/" + imgKey
	return url
}
