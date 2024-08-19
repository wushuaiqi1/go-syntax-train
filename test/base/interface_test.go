package base

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

type Person struct {
}

func NewPerson() *Person {
	return nil
}

func (person Person) Read(p []byte) (n int, err error) {
	return
}

func TestInterfaceNil(t *testing.T) {
	// 注意 go中判空需要值和类型都为nil
	var p io.Reader = NewPerson()
	fmt.Println(p == nil)
	fmt.Println(reflect.ValueOf(p).IsNil())
	fmt.Println(reflect.TypeOf(p))

	// 捕获异常～
	defer func() {
		if ok := recover(); ok != nil {
			fmt.Println(ok)
		}
	}()
	_, err := p.Read([]byte("111"))
	if err != nil {
		return
	}
}
