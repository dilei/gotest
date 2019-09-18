package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	if x := 10; x >= 10 {

	}
	// 子线程painc会导致主进程挂掉
	go say("world")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("hello")

	var i int
	o := make(chan int)
	go func(i int) {
		i = 8
		fmt.Println("i in go", i)
		o <- 1
	}(i)
	i = 9
	<-o
	fmt.Println(i)
}
