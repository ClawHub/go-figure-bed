package dao

import (
	"go-figure-bed/pkg/gmysql"
	"go-figure-bed/pkg/models"
)

//插入
func Insert(api, email, mainUrl, url, imgName, imgMime string) error {
	fb := models.FigureBed{
		Api:     api,
		Email:   email,
		MainUrl: mainUrl,
		Url:     url,
		ImgName: imgName,
		ImgMime: imgMime,
	}
	if err := gmysql.DB.Create(&fb).Error; err != nil {
		return err
	}
	return nil
}
