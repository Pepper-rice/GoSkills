package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	ctx := context.Background()
	// Redis kv结构
	rdb.Set(ctx, "redis:string", "bar", time.Second*10)
	val, _ := rdb.Get(ctx, "redis:string").Result()
	fmt.Println("redis kv 获取的值是:", val)
	rdb.Del(ctx, "redis:string")

	// Redis list结构
	rdb.LPush(ctx, "redis:list", "bar", "foo")
	listResult, _err := rdb.LRange(ctx, "redis:list", 0, -1).Result()
	if _err != nil {
		panic(_err)
	}
	for index, item := range listResult {
		fmt.Println(fmt.Sprintf("redis list 获取的值是: %s,index: %d ", item, index))
	}
	rdb.Del(ctx, "redis:list")

	// Redis hash结构
	rdb.HSet(ctx, "redis:hash", "foo", "bar")
	hashResult, _err := rdb.HGetAll(ctx, "redis:hash").Result()
	if _err != nil {
		panic(_err)
	}
	for key, value := range hashResult {
		fmt.Println(fmt.Sprintf("redis hash 获取的值是: key: %s, value: %s", key, value))
	}
	rdb.Del(ctx, "redis:hash")

	// Redis set结构
	rdb.SAdd(ctx, "redis:set", "foo", "bar")
	setResult, _err := rdb.SMembers(ctx, "redis:set").Result()
	if _err != nil {
		panic(_err)
	}
	for _, item := range setResult {
		fmt.Println(fmt.Sprintf("redis set 获取的值是: %s", item))
	}
	// Redis set结构删除
	rdb.SRem(ctx, "redis:set", "foo")
	setResult, _err = rdb.SMembers(ctx, "redis:set").Result()
	if _err != nil {
		panic(_err)
	}
	for _, item := range setResult {
		fmt.Println(fmt.Sprintf("redis set 删除后获取的值是: %s", item))
	}
	rdb.Del(ctx, "redis:set")
}
