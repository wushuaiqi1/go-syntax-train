package base

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAssert(t *testing.T) {
	// 类型断言
	arr1 := []any{http.MethodGet, Handler}
	ant := getAny(arr1)
	res, ok := ant.([]any)
	if ok {
		fmt.Println(res)
		fmt.Println(reflect.TypeOf(res))
	}
	method, ok := res[0].(string)
	if ok {
		fmt.Println(method)
	}
	execute, ok := res[1].(func())
	if ok {
		execute()
	}
}

func getAny(datasource any) any {
	return datasource
}

func Handler() {
	fmt.Println("handler ...")
}
