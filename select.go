package main

import "time"

func main() {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	var flag1, flag2 bool
	for {

		select {
		case <-ch:
			println("case1")
			flag1 = true
		case <-ch:
			println("case2")
			flag2 = true
		}

		if flag1 && flag2 {
			break
		}
	}
}
