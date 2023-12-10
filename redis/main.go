package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	useRedis()
}

func useRedis() {
	client := redis.NewClient(&redis.Options{
		DB:         0,
		Addr:       "127.0.0.1:6379",
		Password:   "",
		MaxRetries: 5,
	})
	s2, err2 := client.Ping().Result()
	if err2 != nil {
		fmt.Printf("%v", err2)
	}
	fmt.Println(s2)
	result, err := client.Set("zhengyb", 1, 2).Result()
	if err != nil {
		fmt.Errorf("set redis key error:%w", err)
	}
	fmt.Println("result", result)

	s, err := client.Get("zhengyb").Result()
	fmt.Println(s, err)
}
