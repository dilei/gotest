package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 12
	fmt.Println(strconv.FormatInt(int64(num), 10))

	fmt.Println(strconv.FormatInt(1, 2))
}
