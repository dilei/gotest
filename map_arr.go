package main

import "fmt"

func main() {
	mapArr := make(map[int][]int)
	arr := []int{1, 2, 3, 4, 5}

	for k, v := range arr {
		mapArr[k] = append(mapArr[k], v)
	}

	fmt.Println(mapArr)
}
