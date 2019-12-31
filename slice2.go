package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[0:3])
	fmt.Println(arr[3:3])

	arr2 := arr[:]
	arr2[0] = 4

	fmt.Println(arr)
	fmt.Println(arr2)

	test(arr)
	test(arr)
	test(arr[:])
}

func test(arr []int) {
	arr[0] = 10
	fmt.Printf("%p\n", &arr)
	fmt.Println(arr)
}
