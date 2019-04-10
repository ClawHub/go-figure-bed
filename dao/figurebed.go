package dao

import (
	"go-figure-bed/pkg/gmysql"
	"go-figure-bed/pkg/models"
)

//插入
func Insert(api, email, mainUrl, url, imgName, imgMime string) (err error) {
	fb := models.FigureBed{
		Api:     api,
		Email:   email,
		MainUrl: mainUrl,
		Url:     url,
		ImgName: imgName,
		ImgMime: imgMime,
	}
	err = gmysql.DB.Create(&fb).Error
	return
}

//查询所有
func QueryAll(mainUrl, email string) (fbs []models.FigureBed, err error) {
	err = gmysql.DB.Where("email = ? AND main_url = ? AND deleted_on = ? ", email, mainUrl, 0).Find(&fbs).Error
	return
}

//查询一个
func QueryOne(mainUrl, email string) (fb models.FigureBed, err error) {
	err = gmysql.DB.Where("email = ? AND main_url = ? AND deleted_on = ? ", email, mainUrl, 0).First(&fb).Error
	return
}
