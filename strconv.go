package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 12
	fmt.Println(strconv.FormatInt(int64(num), 10))

	fmt.Println(strconv.FormatInt(1, 2))

	// intK, err := strconv.ParseInt("", 10, 64)
	intK, err := strconv.Atoi("")
	// 空字符串转int出错，intK默认为int的零值
	fmt.Println(err.Error())
	fmt.Println(intK)
}
