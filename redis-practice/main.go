package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	ctx := context.Background()

	rdb, err := createRDB(ctx)
	if err != nil {
		fmt.Printf("redis create error : %v\n", err)
	}

	val, err := rdb.Do(ctx, "get", "Sean").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("redis get key not existed")
	case err != nil:
		fmt.Println("redis get key error", err)
	case val == "":
		fmt.Println("redis get key is empty string")
	}

	if err := rdb.Set(ctx, "Sean", "123", 300*time.Second).Err(); err != nil {
		fmt.Printf("redis set value error : %v\n", err)
	}

	val, err = rdb.Get(ctx, "Sean").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("redis get key not existed")
	case err != nil:
		fmt.Println("redis get key error", err)
	case val == "":
		fmt.Println("redis get key is empty string")
	}

	fmt.Printf("redis get Sean : %v\n", val)
}

func createRDB(ctx context.Context) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{ // 建立 client 連線不會拋出錯誤
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	_, err = rdb.Ping(ctx).Result() // 使用 ping 方法確認是否有連線成功
	if err != nil {
		return
	}

	fmt.Println("redis connect success")
	return
}
