package test

import (
	"fmt"
	"math/rand"
	"testing"
)

// 测试压栈
func TestFunc(t *testing.T) {
	data := make([]int, 10)
	dfs(data)
	fmt.Println(data)
}

// 压栈本质上传递的是切片
func dfs(data []int) {
	if len(data) > 2 {
		dfs(data[1:])
	}
	data[0] = rand.Int()
}
