package main

import (
	"fmt"
	"reflect"
)

// 从基本类型获取复合类型
var (
	Int = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

func main() {
	c := reflect.ChanOf(reflect.SendDir, String)
	fmt.Println(c)

	m := reflect.MapOf(String, Int)
	fmt.Println(m)

	s := reflect.SliceOf(Int)
	fmt.Println(s)

	t := struct{ Name string }{}
	p := reflect.PtrTo(reflect.TypeOf(t))
	fmt.Println(p)

	// 复合类型获取基本类型
	t1 := reflect.TypeOf(make(chan int)).Elem()
	fmt.Println(t1)


	s2 := make([]int, 0, 10)
	v := reflect.ValueOf(&s2).Elem()
	v.SetLen(2)
	v.Index(0).SetInt(100)
	v.Index(1).SetInt(200)
	fmt.Println(v.Interface(), s)
	v2 := reflect.Append(v, reflect.ValueOf(300))
	v2 = reflect.AppendSlice(v2, reflect.ValueOf([]int{400, 500}))

	fmt.Println(v2.Interface())
	fmt.Println("----------------------")
	m2 := map[string]int{"a": 1}
	v = reflect.ValueOf(&m2).Elem()
	v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100)) // update
	v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200)) // add
	fmt.Println(v.Interface(), m)
}
