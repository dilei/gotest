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

	str := "Go爱好者123"
	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}

	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))

	fmt.Println(string(-1))
	fmt.Println(string(1))
	fmt.Println(string(3))
	// 以下都是表示的数字 1
	fmt.Println(string(49))   // 49 表示 十进制的49（061   49    31    1）
	fmt.Println(string(0x31)) // 表示 十六进制的49（061   49    31    1）
	fmt.Printf("[% x]\n", []byte(string(49)))
	// 以下都是表示的字符 a
	fmt.Println(string(0x61)) // 表示 十六进制61（141   97    61    a）

	str2 := str
	// 为什么输出不是指向的同一个地址？
	// type stringStruct struct {
	// 	str unsafe.Pointer
	// 	len int
	// }
	// 因为string是一个结构体，结构体内部是指向字符串的指针(str)，str属性指向的是相同的地址，二字符串str2和str不是同一个结构体
	fmt.Printf("str: %p str2: %p\n", &str, &str2) // str: 0xc0000361f0 str2: 0xc000036210
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
