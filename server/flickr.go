package server

import (
	"encoding/xml"
	"go-figure-bed/pkg/e/bed"
	"go-figure-bed/pkg/logging"
	"go-figure-bed/pkg/setting"
	"go.uber.org/zap"
	"io"

	"gopkg.in/masci/flickr.v2"
)

func flickrGetOauth() *flickr.FlickrClient {
	apiKey := setting.BedSetting.Flickr.ApiKey
	apiSecret := setting.BedSetting.Flickr.ApiSecret
	client := flickr.NewFlickrClient(apiKey, apiSecret)
	client.OAuthToken = setting.BedSetting.Flickr.OauthToken
	client.OAuthTokenSecret = setting.BedSetting.Flickr.OauthTokenSecret
	client.Id = setting.BedSetting.Flickr.Id
	client.OAuthSign()
	return client
}

func flickrGetPicInfo(id string) string {
	client := flickrGetOauth()
	client.Init()
	client.Args.Set("method", "flickr.photos.getInfo")
	client.Args.Set("photo_id", id)
	client.OAuthSign()
	response := &flickr.BasicResponse{}
	err := flickr.DoGet(client, response)

	if err != nil {
		logging.AppLogger.Error("Flickr Error:", zap.Error(err))
		return ""
	}
	v := bed.FlickrGetPicResp{}
	_ = xml.Unmarshal([]byte(response.Extra), &v)
	if v.Originalformat != "gif" && setting.BedSetting.Flickr.DefaultSize != "o" {
		picUrl := "https://" + "farm" + v.Farm + ".staticflickr.com/" + v.Server + "/" + v.Id + "_" + v.Secret + "_" + setting.BedSetting.Flickr.DefaultSize + ".jpg"
		return picUrl
	} else {
		picUrl := "https://" + "farm" + v.Farm + ".staticflickr.com/" + v.Server + "/" + v.Id + "_o-" + v.Originalsecret + "_o." + v.Originalformat
		return picUrl
	}
}

//上传雅虎图片服务
func UploadToFlickr(file io.Reader, fileName string) string {
	if !setting.BedSetting.Flickr.OpenFlickrStore {
		return ""
	}
	client := flickrGetOauth()
	resp, err := flickr.UploadReader(client, file, fileName, nil)
	if err != nil {
		logging.AppLogger.Error("Flickr ERROR :", zap.Error(err))
	}
	return flickrGetPicInfo(resp.ID)
}
