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

	str3 := "||23||2|3|3|"
	str3 = strings.Trim(str3, "|")
	strArr := strings.Split(str3, "|")
	fmt.Println(strArr)

	_ = strArr
	fmt.Println(len(strArr))
}
