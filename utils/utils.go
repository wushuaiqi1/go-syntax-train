package utils

import "strings"

func GetUri(method string, pattern string) (res string) {
	return method + "#" + pattern
}

// ResolveRequestPath 解析请求路径
// for example "/user/login/" => ["user","login"]
func ResolveRequestPath(pattern string) []string {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")
	return paths
}
