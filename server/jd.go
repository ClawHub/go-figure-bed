package server

import (
	"go-figure-bed/pkg/utils"
	"math/rand"
	"regexp"
	"strconv"
)

func UploadToJd(img []byte, imgName string, imgType string) string {
	url := "https://search.jd.com/image?op=upload"

	file := &utils.FormFile{
		Name:  imgName,
		Key:   "file",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := utils.FormPost(file, url, header)
	var pre = regexp.MustCompile(`(?m)ERROR`)

	if !pre.MatchString(data) {
		var re = regexp.MustCompile(`(?m)\("(.*)"\)`)
		imgFix := re.FindAllStringSubmatch(data, -1)[0][1]
		url = "https://img" + strconv.Itoa(rand.Intn(3)+11) + ".360buyimg.com/img/" + imgFix
		return url
	} else {
		return ""
	}

}
