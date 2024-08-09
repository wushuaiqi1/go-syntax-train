package test

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	a := 111.1
	b := float32(a)
	c := float64(b)
	fmt.Println(b)
	fmt.Println(c)
}
