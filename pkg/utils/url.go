package utils

//格式化 url
func FormatUrl(url *string) {
	n := len(*url)
	rs := []rune(*url)
	s := string(rs[n-1 : n])
	if s != "/" {
		*url += "/"
	}
	s = string(rs[0:1])
	if s == "/" {
		*url = string(rs[1:n])
	}
}
