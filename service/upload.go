package service

import (
	"go-figure-bed/dao"
	"go-figure-bed/pkg/setting"
	"go-figure-bed/pkg/utils"
	"go-figure-bed/server"
	"mime/multipart"
	"sync"
)

//处理文件上传
func Handle(email string, f multipart.File, h *multipart.FileHeader) (mainUrl string) {
	imgMime := h.Header.Get("Content-Type")
	imgInfo := h.Header.Get("Content-Disposition")
	//读取文件名称
	imgName := utils.GetFileNameByMimeType(imgInfo)
	//读取文件
	size := h.Size
	fileContent := make([]byte, size)
	_, _ = f.Read(fileContent)

	//生成本次图片的URL
	mainUrl = utils.GetVirtualUrl(imgName)

	//根节点处理，使用工作池
	var wg sync.WaitGroup
	for _, api := range setting.AppSetting.RootSiteApi {
		wg.Add(1)
		//协程处理根节点
		go dealRootSite(email, api, imgMime, imgName, mainUrl, fileContent, h, &wg)

	}
	wg.Wait()
	//异步分发
	for _, api := range setting.AppSetting.OtherSiteApi {
		go WhichApi(email, api, imgMime, imgName, mainUrl, fileContent, h)

	}
	return
}

//根节点处理
func dealRootSite(email, api, imgMime, imgName, mainUrl string, fileContent []byte, h *multipart.FileHeader, wg *sync.WaitGroup) {
	WhichApi(email, api, imgMime, imgName, mainUrl, fileContent, h)
	wg.Done()
}

//使用图床api
func WhichApi(email, api, imgMime, imgName, mainUrl string, fileContent []byte, h *multipart.FileHeader) {
	var url string
	//API图床选择
	switch api {
	case "Local":
		showUrl, _, _, _ := server.UpLoadToLocal(imgName, fileContent)
		url = showUrl
	case "SouGou":
		url = server.UpLoadToSouGou(fileContent)
	case "Sina":
		if setting.BedSetting.Sina.OpenSinaPicStore == false {
			url = ""
		} else {
			url = server.UpLoadToSina(fileContent, imgMime)
		}
	case "Smms":
		url = server.UploadToSmms(fileContent, imgName)
	case "CC":
		durl, _ := server.UploadToCC(fileContent, imgName, imgMime)
		url = durl
	case "Flickr": //雅虎图片服务器，国内访问吃力
		if setting.BedSetting.Flickr.OpenFlickrStore == false {
			url = ""
		} else {
			file, err := h.Open()
			if err != nil {
			}
			url = server.UploadToFlickr(file, h.Filename)
		}
	case "Baidu":
		url = server.UploadToBaidu(fileContent, imgName)
	case "Qihoo":
		url = server.UploadToQihoo(fileContent, imgName, imgMime)
	case "NetEasy":
		url = server.UploadToNetEasy(fileContent, imgName, imgMime)
	case "Jd":
		url = server.UploadToJd(fileContent, imgName, imgMime)
	case "JueJin":
		url = server.UploadToJueJin(fileContent, imgName, imgMime)
	case "Ali":
		url = server.UploadToAli(fileContent, imgName, imgMime)
	case "Open":
		url = server.UpLoadToPublicSina(fileContent, imgName, imgMime)
	}
	if url != "" {
		// api  email mainUrl url imgName imgMime 入库
		_ = dao.Insert(api, email, mainUrl, url, imgName, imgMime)
	}
}
