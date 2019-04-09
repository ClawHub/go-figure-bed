package bed

//新浪公共接口，只需要提供 api 地址即可
//{"code":1,"msg":"操作成功","data":{"code":"200","width":176,"height":254,"size":13476,"pid":"005BYqpgly1fz9xxss19rj372jrq","url":"https:\/\/ws3.sinaimg.cn\/large\/005BYqpgly1fz9xxss19rj304w072jrq.jpg"},"runtime":"0.311697s"}
type SinaPublicResponse struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}
