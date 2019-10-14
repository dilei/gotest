package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	fmt.Printf("%v", ch)

	close(ch)

	fmt.Println(<-ch)

	select {
	case <-ch:
		fmt.Println("hello world")
	case <-time.After(time.Second):
		fmt.Println("time out")
	}
}
