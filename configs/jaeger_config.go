package configs

import (
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go/config"
	"log"
)

func LoadJaegerConfig() {
	// 配置 Jaeger 客户端

	tags := make([]opentracing.Tag, 2)
	tag := opentracing.Tag{
		Key: "服务", Value: "Go语法训练服务"}
	tags = append(tags, tag)

	cfg := config.Configuration{
		Tags:        tags,
		ServiceName: viper.GetString("server.name"), // 你的服务名称
		Sampler: &config.SamplerConfig{
			Type:  viper.GetString("tracing.sampler.type"),   // 采样类型，这里使用固定采样
			Param: viper.GetFloat64("tracing.sampler.param"), // 采样比例，1表示采样所有请求
		},
		Reporter: &config.ReporterConfig{
			LogSpans:          true,                                // 日志记录 Span 信息
			CollectorEndpoint: viper.GetString("tracing.endpoint"), // Jaeger HTTP 接口的地址
		},
	}

	// 初始化 Jaeger tracer
	tracer, _, err := cfg.NewTracer()
	if err != nil {
		log.Printf("could not create Jaeger tracer: %v", err)
	}

	// 将 tracer 设置为全局 tracer
	opentracing.SetGlobalTracer(tracer)
	log.Println("load jaeger config success...")
}
