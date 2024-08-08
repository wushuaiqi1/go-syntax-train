package v2

import "log"

// Node 多叉树节点
type Node struct {
	children   []*Node
	pattern    string
	datasource any
}

func NewNode() *Node {
	return &Node{}
}

func newNode(pattern string) *Node {
	return &Node{
		children: make([]*Node, 0, 8),
		pattern:  pattern,
	}
}

// SearchNode 根据所给路径找到目标节点
func (root *Node) SearchNode(paths []string) any {
	// 根节点不存在子节点
	if root.children == nil || len(root.children) == 0 {
		return nil
	}
	// 层序遍历搜素
	cur := root
	for i := 0; i < len(paths); i++ {
		path := paths[i]
		// 存在路径节点
		if find, ok := cur.findMatchChild(path); ok {
			cur = find
			continue
		}
		// 不存在路径节点
		if i != len(paths)-1 {
			return nil
		}
	}
	// 返回目标节点
	return cur.datasource
}

// CreateNode 创建节点
func (root *Node) CreateNode(paths []string, data any) {
	cur := root
	for idx, path := range paths {
		if find, ok := cur.findMatchChild(path); ok {
			// 1.命中子节点，继续寻找
			cur = find
		} else {
			// 2.未命中，创建子树
			cur.createSubTree(paths[idx:], data)
			return
		}
	}
}

// findMatchChild 孩子节点是否存在该路径
func (root *Node) findMatchChild(path string) (find *Node, ok bool) {
	for _, child := range root.children {
		if child.pattern == path {
			return child, true
		}
	}
	return
}

// createSubTree 创建子树
func (root *Node) createSubTree(paths []string, data any) {
	cur := root
	for _, path := range paths {
		node := newNode(path)
		cur.children = append(cur.children, node)
		cur = node
	}
	// 到达叶子节点,添加数据源
	cur.datasource = data
}

// MiddleTraversal 中序遍历
func (root *Node) MiddleTraversal() {
	log.Println(root.pattern)
	for i := 0; i < len(root.children); i++ {
		// 当前孩子节点
		child := root.children[i]
		child.MiddleTraversal()
	}
}
