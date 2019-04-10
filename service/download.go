package service

import (
	"go-figure-bed/dao"
)

//下载
func Download(mainUrl, email string) (url string) {
	fb, err := dao.QueryOne(mainUrl, email)
	if err != nil {
		return ""
	}
	return fb.Url
}
