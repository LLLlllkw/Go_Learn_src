package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
)

const N = 100000 // 插入十万条数据

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   1,
	})

	ClearDB(ctx, client)
	pipeline := client.Pipeline()
	for i := range [N]struct{}{} {

		pipeline.Set(ctx, fmt.Sprintf("hello:%d", i), GenerateValueWithBytes(100), 0)
	}
	if _, err := pipeline.Exec(ctx); err != nil {
		panic(err)
	}

	fmt.Println("done")
}

func GenerateValueWithBytes(size int) string {
	var res []byte
	str := "ABCDEFGHIJKLMNOPQRSTUVWSYZ0123456789"
	for i := 0; i < size; i++ {
		res = append(res, str[rand.Intn(len(str))])
	}
	return string(res)
}


func ClearDB(ctx context.Context, client *redis.Client) {
	_ = client.FlushDB(ctx)