package service

import "github.com/opentracing/opentracing-go"

// TraceMark 数据埋点
func TraceMark(spanName string, options map[any]any) {
	curSpan := opentracing.StartSpan(spanName)
	defer curSpan.Finish()
	for key, value := range options {
		curSpan.LogKV(key, value)
	}
}
