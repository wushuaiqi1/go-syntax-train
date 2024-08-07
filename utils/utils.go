package utils

func GetUri(method string, pattern string) (res string) {
	return method + "#" + pattern
}
