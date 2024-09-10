package controller

import (
	"github.com/uber/jaeger-client-go/config"
	"go-syntax-train/configs"
	"net/http"
)

func GetUserInfo() http.HandlerFunc {
	configs.Tracer.StartSpan("获取用户信息")
	defer config.Extractor()
}
