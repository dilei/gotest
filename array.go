package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr2 := arr[0:2]
	fmt.Println(arr2)

	var a interface{} = "string"
	str, ok := a.(string)
	if ok {
		fmt.Println(str)
	}
}
