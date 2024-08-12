package tree

import (
	"fmt"
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

func List() {
	for _, val := range list {
		fmt.Println(val)
	}
}

// GetPathList 获取二叉树所有路径
func GetPathList(root *TreeNode, data []int) {
	// 当前元素添加到切片
	data = append(data, root.Val)
	// 叶子节点处理
	if root.Left == nil && root.Right == nil {
		// 切片底层共享数据
		path := make([]int, len(data))
		copy(path, data)
		fmt.Printf("%p\n", data)
		list = append(list, data)
		return
	}
	if root.Left != nil {
		GetPathList(root.Left, data)
	}
	if root.Right != nil {
		GetPathList(root.Right, data)
	}
}

// convertBiNode 二叉搜索树转换为单向链表
func convertBiNode(root *TreeNode) *TreeNode {
	head := &TreeNode{-1, nil, root}
	inorderConvert(head, root)
	return head.Right
}

func inorderConvert(prev *TreeNode, cur *TreeNode) *TreeNode {
	if cur == nil {
		return prev
	}
	prev = inorderConvert(prev, cur.Left)
	// 父亲节点断开左孩子
	cur.Left = nil
	// 前驱节点孩子节点指向父亲节点
	prev.Right = cur
	// 前驱节点转移到父节点
	prev = cur
	prev = inorderConvert(prev, cur.Right)
	return prev
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(depth(root.Left), depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	return max(depth(cur.Left), depth(cur.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

var res []int

func InorderTraversal(root *TreeNode) []int {
	return inorder(root, make([]int, 0))
}

// inorder 中序遍历 注意数组返回上层
func inorder(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	if root.Left != nil {
		res = inorder(root.Left, res)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = inorder(root.Right, res)
	}
	return res
}

// isSameTree 相同的树
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

// isSymmetric 对称的树
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return check(p.Left, q.Right) && check(p.Right, q.Left)
	} else {
		return false
	}
}

func sortedArrayToBST(nums []int) (res *TreeNode) {
	return buildTree(nums, 0, len(nums)-1)
}
func buildTree(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	cur := &TreeNode{
		Val: nums[mid],
	}
	cur.Left = buildTree(nums, left, mid-1)
	cur.Right = buildTree(nums, mid+1, right)
	return cur
}
