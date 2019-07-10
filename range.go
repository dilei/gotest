package main

import "fmt"

func main() {
	str := "中国"
	for j :=0; j <2; j++ {
		fmt.Println(str[j])  // utf-8 byte
	}

	for key := range str {
		fmt.Println(key) // unicode rune
	}

	for _, val := range str {
		fmt.Println(val)  // rune
	}

	i := 0
	for range str { // 只遍历
		i++
		fmt.Println(i)
	}
}
