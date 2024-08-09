package algo

import (
	"go-syntax-train/utils"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// decorateRecord 层序遍历
func decorateRecord(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		// 更新访问长度
		curLen := len(queue)
		curLevel := make([]int, 0)
		// 处理当前层元素
		for idx := 0; idx < curLen; idx++ {
			curNode := queue[idx]
			curLevel = append(curLevel, curNode.Val)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		// 移除上一层元素
		queue = queue[curLen:]

		res = append(res, curLevel)
	}
	return res
}

func sumRootToLeaf(root *TreeNode) int {
	GetPathList(root, make([]int, 0))
	// 求和
	count := 0
	for _, numbers := range list {
		count += utils.BinarySum(numbers)
	}
	return count
}

var list [][]int = make([][]int, 0)

func GetPathList(root *TreeNode, data []int) {
	// 当前元素添加到切片
	data = append(data, root.Val)
	// 叶子节点处理
	if root.Left == nil && root.Right == nil {
		// 切片底层共享数据
		path := make([]int, len(data))
		copy(path, data)
		list = append(list, path)
		return
	}
	if root.Left != nil {
		GetPathList(root.Left, data)
	}
	if root.Right != nil {
		GetPathList(root.Right, data)
	}
}
