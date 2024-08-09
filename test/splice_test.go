package test

import (
	"fmt"
	"go-syntax-train/algo"
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

func TestInit(t *testing.T) {
	stack := make([]int, 0)
	stack = append(stack, 1, 3, 4, 4)
	fmt.Println(stack)
}

func TestCommon(t *testing.T) {
	// 共享切片
	data := make([]int, 0)
	fmt.Printf("%p\n", data)
	test1(data)
	fmt.Println(data)
}
func test1(data []int) {
	// 新地址
	data = append(data, 1)
	fmt.Printf("%p\n", data)
	fmt.Println(data)
}

func TestGetPathList(t *testing.T) {
	tree := &algo.TreeNode{
		Val: 0,
		Left: &algo.TreeNode{
			Val: 0,
			Left: &algo.TreeNode{
				Val: 1,
			},
			Right: &algo.TreeNode{
				Val: 0,
			},
		},
		Right: &algo.TreeNode{
			Val: 1,
			Left: &algo.TreeNode{
				Val: 1,
			},
			Right: &algo.TreeNode{
				Val: 1,
				Left: &algo.TreeNode{
					Val: 0,
				},
				Right: &algo.TreeNode{
					Val: 1,
				},
			},
		},
	}
	algo.GetPathList(tree, make([]int, 0))
}
