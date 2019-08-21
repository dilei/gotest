package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "123"
	str2 := "123"

	if str > str2 {

	}

	strArr := strings.Split(str, ",")
	fmt.Println(strArr)

	_ = strArr
	fmt.Println(strArr)
}
