package base

import (
	"context"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	ctx := context.Background()
	res := ctx.Value("111")
	println(res)
}

func TestContextTODO(t *testing.T) {

}
