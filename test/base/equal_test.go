package base

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	a := make([]int, 1)
	fmt.Printf("%p\n", &a[0])
	a = append(a, 10)
	fmt.Printf("%p\n", &a[0])
	fmt.Println(len(a), cap(a))
	a = append(a, 2, 3, 4)
	fmt.Printf("%p\n", &a[0])
	//fmt.Printf("%p\n", &a)
	//fmt.Println(&a == &b)
}
