package main

import "fmt"

func main() {
	str := "123"
	var sstr s
	sstr = s(str)
	fmt.Println(sstr)

	var str2 ss
	str2 = str
	fmt.Println(str2)
}

type s string
type ss = string
