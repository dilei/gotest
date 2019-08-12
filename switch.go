package main

import (
	"fmt"
	"time"
)

func main() {
	arr64 := []int64{1, 2, 3, 4}
	for _, val := range arr64 {
		switch val {
		case 1:
			fmt.Println(1)
			fmt.Println(time.Time{})
			fallthrough
		default:
			fmt.Println("default")
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println(3)
		}
	}
}
