package v2

type RouterHandler interface {
	Register()
	GetFuncByRoute()
}

type RouteHandlerBasedOnTree struct {
	root *Node
}
