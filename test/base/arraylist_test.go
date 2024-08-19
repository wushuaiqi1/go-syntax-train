package base

import (
	go_syntax_train "go-syntax-train/base"
	"testing"
)

func TestAdd(t *testing.T) {
	arrayList := go_syntax_train.NewArrayList()
	arrayList.Add(1000)
	arrayList.Add("hello", "张三", 102)
	arrayList.Show()
}

func TestDelete(t *testing.T) {
	cur := go_syntax_train.NewArrayList()
	cur.Add(100, 20, 30, 400, 50)
	cur.Remove(-1)
	cur.Show()
}
