package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplice(t *testing.T) {
	parent := []int{1, 2, 3, 4, 5}
	fmt.Println(parent[:3])  // [0,3)
	fmt.Println(parent[1:])  // 1～末位
	fmt.Println(parent[1:3]) // [1,3)
	sub1 := parent[1:3]
	// 操作的是原地址
	sub1[1] = -1
	fmt.Println(parent[:3])
	fmt.Println(parent[1:]) // 1～末位
	fmt.Println(sub1)
}

func TestAppend(t *testing.T) {
	var array [10]int
	splice := array[:]
	// 数组
	fmt.Println(reflect.TypeOf(array))
	// 切片
	fmt.Println(reflect.TypeOf(splice))
	splice = append(splice, 1, 0, 0, 0, 1, 2, 0, 1, 1, 0, 1, 1, 1, 1, 1111, 1, 1, 1, 1)
	fmt.Println(splice)
	fmt.Println(len(splice))
	fmt.Println(cap(splice))
}
