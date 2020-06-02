package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(string(-1))

	var arr []string
	arr2 := []int{1, 2, 3, 4}
	fmt.Println(arr, arr2)
	fmt.Println(len(arr), len(arr2))

	arr3 := append(([]int)(nil), arr2...)

	var arr5 []int
	arr4 := append(arr5, arr2...)
	fmt.Println(arr3)
	fmt.Println(arr4)

	fmt.Println(strings.Compare("24.06", "9"))

	fmt.Printf("%v", numDecodings("110"))
}

func clean(str string) (string, string) {
	return str, str[:]
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	f := make([]int, len(s)+1)
	f[0] = 1
	// 从1~n依次计算f[i]
	for i := 1; i <= len(s); i++ {
		// 枚举最后一个字母匹配的位数
		if s[i-1] > '0' {
			f[i] += f[i-1]
		}
		if i > 1 && s[i-2:i] >= "10" && s[i-2:i] <= "26" {
			f[i] += f[i-2]
		}
	}
	// 返回答案
	return f[len(s)]
}
