package main

import (
    "bytes"
    "fmt"
	"math"
    "strings"
    "unicode/utf8"
)

func main() {
	var f float64
	f = 16777216213243
	fmt.Println(f)

	fmt.Println(math.IsNaN(-1))
	fmt.Println(math.NaN())

	str := "北京100"
	fmt.Println(len(str))
	fmt.Println(f1(str))
	fmt.Println(f2(str))
	fmt.Println(f3(str))
	fmt.Println(f4(str))
	n := 0
	for range str {
		n++
	}
	fmt.Println(n)

	for i := 0 ; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Println("%d\t%c\n", i, r)
		i += size
	}
}

func f1(s string) int {
    return bytes.Count([]byte(s), nil) - 1
}

func f2(s string) int {
    return strings.Count(s, "") - 1
}

func f3(s string) int {
    return len([]rune(s))
}

func f4(s string) int {
    return utf8.RuneCountInString(s)
}
