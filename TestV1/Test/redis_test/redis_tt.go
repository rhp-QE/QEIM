package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 如果有密码，填写密码
		DB:       0,  // 选择数据库
	})
	ctx := context.Background()
	// 设置键值对
	setTimeBegin := time.Now()
	err := client.Set(ctx, "keytest1", "value", 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	setTimeEnd := time.Now()

	fmt.Println(setTimeBegin)
	fmt.Println(setTimeEnd)

	fmt.Println(setTimeEnd.Sub(setTimeBegin))

	setTimeBegin = time.Now()
	for i := 0; i < 100; i++ {
		fmt.Println("fadsfasdf")
	}
	setTimeEnd = time.Now()
	fmt.Println(setTimeEnd.Sub(setTimeBegin))

	// 同步
	value, err := client.Get(ctx, "keytest1").Result()

	// 异步
	var ch = make(chan string)
	go func() {
		val, _ := client.Get(ctx, "keytest1").Result()
		ch <- val
	}()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("key:", value)
}


A
