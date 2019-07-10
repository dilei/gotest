package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type Person1 struct {
	Name string
	Age  int
}

func Round(f float64, prec int, round bool) float64 {
	pow10_n := math.Pow10(prec)
	if round {
		return math.Trunc((f + 0.5/pow10_n)*pow10_n)/pow10_n
	}
	return math.Trunc((f)*pow10_n) / pow10_n
}

func IntSliceDiff(a, b []int) []int {
	diffStr := []int{}
	m := map[int]int{}

	for _, s1Val := range a {
		m[s1Val] = 1
	}
	for _, s2Val := range b {
		m[s2Val] = m[s2Val] + 1
	}

	for mKey, mVal := range m {
		if mVal==1 {
			diffStr = append(diffStr, mKey)
		}
	}

	return diffStr
}

func testRes() (res int) {
	var tmp = 1000
	return tmp
}

func main() {
	/*
	var pmap = make(map[int]Person1)

	var p1 = Person1{"aaa", 12}
	var p2 = Person1{"bbb", 13}

	pmap[1] = p1
	pmap[2] = p2

	for _, v := range pmap {
		v.Name = "a"
	}

	fmt.Println(pmap)

	for k, v := range pmap {
		v.Name = "a"
		pmap[k] = v
	}
	fmt.Println(pmap)
	*/

	// f  := Round(2.56, 3, true)
	// fmt.Println(f)

	intArr1 := []int{1,2,3}
	intArr2 := []int{2,3,4}
	intArr3 := IntSliceDiff(intArr1, intArr2)
	fmt.Println(intArr3)

	fmt.Println(len([]int{}))

	tmp := []int{}
	for k, v := range tmp {
		fmt.Println(k, v)
	}

	map1 := make(map[int]map[string]int)
	fmt.Println(len(map1))

	fmt.Println(testRes())

	// 接口类型将输出其内部包含的值
	var i interface{} = struct {
		name string
		age  int
	}{"AAA", 20}
	fmt.Printf("%v\n", i)  // 只输出字段值
	fmt.Printf("%+v\n", i) // 同时输出字段名
	fmt.Printf("%#v\n", i) // Go 语法格式

	// 带引号字符串
	s1 := "Hello 世界!"       // CanBackquote
	s2 := "Hello\n世界!"      // !CanBackquote
	fmt.Printf("%q\n", s1)  // 双引号
	fmt.Printf("%#q\n", s1) // 反引号成功
	fmt.Printf("%#q\n", s2) // 反引号失败
	fmt.Printf("%+q\n", s2) // 仅包含 ASCII 字符

	// Unicode 码点
	fmt.Printf("%U, %#U\n", '好', '好')
	fmt.Printf("%U, %#U\n", '\n', '\n')

	str := "颜色_蓝色_浅蓝色:尺码_S_S"
	fmt.Println(str[0:6])

	str = "12,13,14"
	fmt.Println(strings.Split(str, ":"))

	fmt.Println(time.Duration(1) * time.Minute)

	tp, _ := time.ParseDuration("15.5s")
	fmt.Println(tp.Truncate(1000000), tp.Seconds(), tp.Nanoseconds())
	res := tp.Truncate(3*time.Second).Seconds()
	fmt.Println(res, tp.Truncate(1000), tp.Seconds(), tp.Nanoseconds())

	sss := 15 % 10
	fmt.Println(sss)

	arr := []int{1,2,3}
	for key := range arr {
		fmt.Println(key)
	}



}

