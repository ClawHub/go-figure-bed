package utils

import "encoding/base64"

//解密
func Decode(enc *base64.Encoding, str string) string {
	data, err := enc.DecodeString(str)

	if err != nil {
		panic(err)
	}
	return string(data)
}

//加密
func Encode(enc *base64.Encoding, str string) string {
	bData := []byte(str)
	data := enc.EncodeToString(bData)
	return string(data)
}
