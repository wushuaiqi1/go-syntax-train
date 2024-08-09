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

// BinarySum 二进制求和
func BinarySum(numbers []int) (count int) {
	tmp := 1
	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] == 1 {
			count += tmp
		}
		tmp *= 2
	}
	return count
}
