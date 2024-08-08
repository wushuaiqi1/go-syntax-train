package v2

import (
	"fmt"
	"go-syntax-train/utils"
	"log"
)

type RouterHandler interface {
	RegisterRoute(method string, pattern string, handleFunc func(ctx *Context))
	GetFuncByRoute(method string, pattern string) func(ctx *Context)
}

type RouteHandlerBasedOnTree struct {
	root *Node
}

func NewRouteHandlerBasedOnTree() RouterHandler {
	return &RouteHandlerBasedOnTree{
		root: NewNode(),
	}
}

// RegisterRoute 注册路由
func (h *RouteHandlerBasedOnTree) RegisterRoute(method string, pattern string, handle func(ctx *Context)) {
	paths := utils.ResolveRequestPath(pattern)
	h.root.CreateNode(paths, []any{method, handle})
}

// GetFuncByRoute 获取业务方法
func (h *RouteHandlerBasedOnTree) GetFuncByRoute(method string, pattern string) (handle func(ctx *Context)) {
	log.Printf("GetFuncByRoute req:{%s},{%s}", method, pattern)
	paths := utils.ResolveRequestPath(pattern)
	data := h.root.SearchNode(paths)
	if data == nil {
		return nil
	}
	methodType, handle, err := resolveRouteNodeData(data)
	if err != nil {
		log.Println(err)
		return nil
	}
	if methodType != method {
		return nil
	}
	return handle
}

// resolveRouteNodeData 解析路由节点数据
func resolveRouteNodeData(datasource any) (methodType string, method func(ctx *Context), err error) {
	if datasource == nil {
		err = fmt.Errorf("datasource is nil")
		return
	}
	// 断言数据源是any数组
	sources, ok := datasource.([]any)
	if !ok {
		err = fmt.Errorf("datasource not []any")
		return
	}
	// 数组判断长度
	ok = len(sources) == 2
	if !ok {
		err = fmt.Errorf("the splice length of sources not equals 2")
		return
	}
	// 断言数组索引0位置是字符串
	methodType, ok = sources[0].(string)
	if !ok {
		err = fmt.Errorf("the splice index 0 not is string")
		return
	}
	// 断言数组索引1位置是方法
	method, ok = sources[1].(func(ctx *Context))
	if !ok {
		err = fmt.Errorf("the splice index 1 not is func()")
		return
	}
	return methodType, method, nil
}
