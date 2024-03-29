package main

import (
	"context"
	"fmt"
	"time"
)

func v3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("v3 Done : ", ctx.Err())
			return
		default:
			fmt.Println(ctx.Value("key"))
			time.Sleep(3 * time.Second)
		}
	}
}
func v2(ctx context.Context) {
	fmt.Println(ctx.Value("key"))
	fmt.Println(ctx.Value("v1"))
	// 相同键,值覆盖
	ctx = context.WithValue(ctx, "key", "modify from v2")
	go v3(ctx)
}
func v1(ctx context.Context) {
	if v := ctx.Value("key"); v != nil {
		fmt.Println("key = ", v)
	}
	ctx = context.WithValue(ctx, "v1", "value of v1 func")
	go v2(ctx)
	for {
		select {
		default:
			fmt.Println("print v1")
			time.Sleep(time.Second * 2)
		case <-ctx.Done():
			fmt.Println("v1 Done : ", ctx.Err())
			return
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// 向context中传递值
	ctx = context.WithValue(ctx, "key", "main")
	go v1(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
