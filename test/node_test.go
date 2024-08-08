package test

import (
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
