package v1

import (
	"fmt"
	"time"
)

type Filter func(ctx *Context)

type FilterBuilder func(next Filter) Filter

var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(ctx *Context) {
		start := time.Now().Nanosecond()
		next(ctx)
		end := time.Now().Nanosecond()
		fmt.Printf("用了 %d 纳秒", end-start)
	}
}
