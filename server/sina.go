package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/cache"
	"go-figure-bed/pkg/e/bed"
	"go-figure-bed/pkg/logging"
	"go-figure-bed/pkg/setting"
	"go-figure-bed/pkg/utils"
	"go.uber.org/zap"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

//缓存
var memoryCache, _ = cache.NewCache("memory", `{"interval":3600}`)

//新浪图床登录
func Login(name string, pass string) interface{} {
	url := "https://login.sina.com.cn/sso/login.php?client=ssologin.js(v1.4.15)&_=1403138799543"
	userInfo := make(map[string]string)
	userInfo["UserName"] = utils.Encode(base64.StdEncoding, name)
	userInfo["PassWord"] = pass
	cookie := getCookies(url, userInfo)
	return cookie
}

//获取新浪图床 Cookie
func getCookies(durl string, data map[string]string) interface{} {
	//尝试从缓存里面获取 Cookie
	if memoryCache.Get("SinaCookies") != nil {
		return memoryCache.Get("SinaCookies")
	}
	postData := make(url.Values)
	postData["entry"] = []string{"sso"}
	postData["gateway"] = []string{"1"}
	postData["from"] = []string{"null"}
	postData["savestate"] = []string{"30"}
	postData["uAddicket"] = []string{"0"}
	postData["pagerefer"] = []string{""}
	postData["vsnf"] = []string{"1"}
	postData["su"] = []string{data["UserName"]} //UserName
	postData["service"] = []string{"sso"}
	postData["sp"] = []string{data["PassWord"]} //PassWord
	postData["sr"] = []string{"1920*1080"}
	postData["encoding"] = []string{"UTF-8"}
	postData["cdult"] = []string{"3"}
	postData["domain"] = []string{"sina.com.cn"}
	postData["prelt"] = []string{"0"}
	postData["returntype"] = []string{"TEXT"}
	client := &http.Client{}
	request, err := http.NewRequest("POST", durl, strings.NewReader(postData.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(request)
	defer resp.Body.Close()
	cookie := resp.Cookies()
	//缓存 Cookie 缓存一个小时
	_ = memoryCache.Put("SinaCookies", cookie, time.Second*3600)
	return cookie
}

//上传图片
func UpLoadToSina(img []byte, imgMime string) string {
	//是否开启新浪图床
	if setting.BedSetting.Sina.OpenSinaPicStore == false {
		return ""
	}
	durl := "http://picupload.service.weibo.com/interface/pic_upload.php" +
		"?ori=1&mime=image%2Fjpeg&data=base64&url=0&markpos=1&logo=&nick=0&marks=1&app=miniblog"
	imgStr := base64.StdEncoding.EncodeToString(img)
	//构造 http 请求
	postData := make(url.Values)
	postData["b64_data"] = []string{imgStr}
	client := &http.Client{}
	request, err := http.NewRequest("POST", durl, strings.NewReader(postData.Encode()))
	if err != nil {
		logging.AppLogger.Error("UpLoad To Sina fail", zap.Error(err))
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//设置 cookie
	uncooikes := Login(setting.BedSetting.Sina.UserName, setting.BedSetting.Sina.PassWord)
	//需要进行断言转换
	cookies, ok := uncooikes.([]*http.Cookie)
	if !ok {
		panic(ok)
	}
	for _, value := range cookies {
		request.AddCookie(value)
	}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return getSinaUrl(body, imgMime)
}

//获取 Sina 图床 URL
func getSinaUrl(body []byte, imgMime string) string {
	str := string(body)
	//正则获取
	pat := "({.*)"
	check := "[a-zA-Z0-9]{32}"
	res := regexp.MustCompile(pat)
	rule := regexp.MustCompile(check)
	jsons := res.FindAllStringSubmatch(str, -1)
	msg := bed.SinaMsg{}
	//解析 json 到 struct
	err := json.Unmarshal([]byte(jsons[0][1]), &msg)
	if err != nil {
		logging.AppLogger.Error("get Sina Url fail", zap.Error(err))
	}
	//验证 pid 的合法性
	pid := msg.Data.Pics.Pic_1.Pid
	if rule.MatchString(pid) {
		sinaNumber := fmt.Sprint((crc32.ChecksumIEEE([]byte(pid)) & 3) + 1)
		//从配置文件中获取
		size := setting.BedSetting.Sina.DefultPicSize
		n := len(imgMime)
		rs := []rune(imgMime)
		suffix := string(rs[6:n])
		if suffix != "gif" {
			suffix = "jpg"
		}
		sinaUrl := "https://ws" + sinaNumber + ".sinaimg.cn/" + size + "/" + pid + "." + suffix
		return sinaUrl
	}
	return ""
}
