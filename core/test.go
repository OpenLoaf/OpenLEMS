package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/gcache"
)

func main() {
	c := gcache.New()
	ctx1 := context.Background()
	ctx2 := context.Background()

	_ = c.Set(ctx1, 1, 10, time.Hour)
	_ = c.Set(ctx2, 1, 20, time.Hour)

	v1, _ := c.Get(ctx1, 1)
	fmt.Printf("ctx1 1: %v\n", v1)
}
