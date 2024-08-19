package base

import (
	"go-syntax-train/algo/tree"
	v2 "go-syntax-train/v2"
	"testing"
)

func TestNode(t *testing.T) {
	node := v2.NewNode()
	node.CreateNode([]string{"a", "b", "c"}, "hei")
	node.CreateNode([]string{"a", "b", "e"}, "ss")
	node.CreateNode([]string{"c", "d"}, "sss")
	node.MiddleTraversal()
}

func TestInorder(t *testing.T) {
	node := &tree.TreeNode{
		Left: nil,
		Right: &tree.TreeNode{
			Left: &tree.TreeNode{
				Val: 3,
			},
			Val: 2,
		},
		Val: 1,
	}
	tree.InorderTraversal(node)

}
