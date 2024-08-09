package test

import (
	"fmt"
	"go-syntax-train/utils"
	"testing"
)

func TestBinarySum(t *testing.T) {
	count := utils.BinarySum([]int{0, 0, 1, 1})
	fmt.Println(count)
	fmt.Println(utils.BinarySum([]int{0, 0, 1, 1, 1}))
	fmt.Println(utils.BinarySum([]int{0, 0, 1}))
}
