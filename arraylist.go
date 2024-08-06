package go_syntax_train

import "fmt"

// ArrayList 定义结构体动态数组
type ArrayList struct {
	source []any
}

// NewArrayList 创建动态数组
func NewArrayList() *ArrayList {
	return &ArrayList{}
}

func (array *ArrayList) Add(values ...any) {
	for _, value := range values {
		array.source = append(array.source, value)
	}
}

func (array *ArrayList) Remove(idx int) {
	if idx >= len(array.source) || idx < 0 {
		panic("index out range")
	}
	var tmp []any = array.source[idx+1:]
	array.source = array.source[0:idx]
	for _, value := range tmp {
		array.source = append(array.source, value)
	}
}

func (array *ArrayList) Show() {
	for idx, val := range array.source {
		fmt.Println(idx, val)
	}
}
