package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	ERROR_FILE_IS_EMPTY:       "上传文件为空",
	ERROR_FILE_IS_TOO_LARGE:   "上传文件太大",
	ERROR_FILE_TYPE:           "文件类型错误",
	ERROR_CAN_NOT_GET_IMG_URL: "无法获取第三方图床 URL",
	ERROR_TOO_MANY_IMAGES:     "上传图片太多",
	ERROR_FILE_UPLOAD:         "文件上传失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
