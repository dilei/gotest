package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type prod struct {
	Username string
	age int
}

func main() {
	var p prod
	memGet("123", p)
}

func memGet(key string, t interface{}) {
	tmp := reflect.TypeOf(t)
	fmt.Println(tmp)

	// 反射出来的类型 不能作为基本类型使用
	str := `{"Username":"dilei", "age":12}`
	json.Unmarshal([]byte(str), &tmp)
	fmt.Println(tmp)
}
