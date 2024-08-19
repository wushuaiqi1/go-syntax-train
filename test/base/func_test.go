package base

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	// 1.20版本存在延迟绑定的问题，闭包内部维护外部变量的地址。
	// 1.22版本修复。
	fs := make([]func(), 0)
	for i := 0; i < 10; i++ {
		fs = append(fs, func() {
			fmt.Println(i)
		})
	}
	for _, val := range fs {
		val()
	}

}
