package test

import (
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestPointer(t *testing.T) {
	// 整数类型指针
	var a *int

	// 字符串类型指针
	var b *string

	// 字节类型指针
	var c *byte

	log.Println(a == nil)
	log.Println(b == nil)
	log.Println(c == nil)
	log.Println("=========", reflect.TypeOf(a))

	// 数组类型
	var arrayA []int
	var arrayB *[]int
	log.Println(arrayA == nil)
	log.Println(arrayB == nil)

	// 数组添加元素
	arrayA = append(arrayA, 1)
	arrayB = &arrayA
	log.Println("arrayA ====", arrayA)
	log.Println("arrayB ====", *arrayB)
	log.Println("arrayB index 0 ====", (*arrayB)[0])

}
