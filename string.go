package main

import "fmt"

func main() {
	var arr []string
	arr2 := []int{1, 2, 3, 4}
	fmt.Println(arr, arr2)
	fmt.Println(len(arr), len(arr2))

	arr3 := append(([]int)(nil), arr2...)

	var arr5 []int
	arr4 := append(arr5, arr2...)
	fmt.Println(arr3)
	fmt.Println(arr4)
}
