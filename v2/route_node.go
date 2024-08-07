package v2

import (
	"fmt"
)

// Node 多叉树节点
type Node struct {
	children []Node
	function func(ctx *Context)
	pattern  string
}

func NewNode(children []Node, pattern string) Node {
	return Node{
		children: children,
		function: nil,
		pattern:  pattern,
	}
}

func (cur *Node) AddNode(pattern string) {

}

func (cur *Node) Show() {
	// 递归打印
	fmt.Printf(cur.pattern)
	// 打印到叶子节点换行
	if cur.children == nil {
		fmt.Println()
	}
	for i := 0; i < len(cur.children); i++ {
		cur.children[i].Show()
	}
}
