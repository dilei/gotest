package main

import (
	"fmt"
	"time"
)

func say(s string) {
	panic(123)
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 子线程painc会导致主进程挂掉
	go say("world")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("hello")
}
