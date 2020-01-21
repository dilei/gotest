package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v7"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:        "10.255.242.83:6379",
		Password:    "", // no password set
		DB:          0,  // use default DB
		PoolTimeout: 10 * time.Second,
		PoolSize:    100,
	})
	client.PoolStats()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("testkey", 0, 30*time.Second).Err()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		// go client.Incr("testkey")
		wg.Add(1)
		go Incr(&wg, client)
	}
	wg.Wait()

	val, err := client.Get("testkey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val, err)
}

func Incr(wg *sync.WaitGroup, client *redis.Client) {
	defer wg.Done()
	_, err := client.Incr("testkey").Result()
	if err != nil {
		fmt.Println(err.Error())
	}
}
