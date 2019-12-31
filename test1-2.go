package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Bird struct {
	Name string
	Age  int
}

type Bird2 struct {
	Name  string
	Age   int
	Email map[int]string
	BB    *Bird
}

func (b *Bird) Fly() {
	fmt.Println("flying....")
}

func main() {
	sparrow := &Bird{"Sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()

	typeOft := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOft.Field(i).Name, f.Type(), f.Interface())
	}

	res := resMap()
	fmt.Println("----")
	fmt.Printf("%p", res)
	res = nil
	if _, ok := res[123]; ok {
		fmt.Println(123)
	}

	var val []int
	fmt.Println(val)

	var bird *Bird
	fmt.Println(bird)
	b := Bird{"Sparrow", 3}
	bytes, _ := json.Marshal(b)
	fmt.Println(string(bytes))
	err := json.Unmarshal([]byte(`{"Name":"Sparrow","Age":3}`), &bird)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(reflect.TypeOf(bird))

	varMap := map[int]string{
		1: "123",
		2: "234",
	}
	bytes2, _ := json.Marshal(varMap)
	fmt.Println(string(bytes2))
	var tmp map[int]string
	fmt.Println(tmp)
	tmp2 := make(map[int]string)
	fmt.Println(tmp2)
	err = json.Unmarshal([]byte(`{"1":"123","2":"234"}`), &tmp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(tmp)
	fmt.Println(reflect.TypeOf(tmp))

	var intval []int
	fmt.Println(len(intval))
	intval = append(intval, 0)
	fmt.Println(len(intval))

	var b2 Bird2
	for k, v := range b2.Email {
		fmt.Println(k, v)
	}
	fmt.Println(b2.BB)
	fmt.Println("end")
}

func resMap() map[int]int {
	res := make(map[int]int)
	res[1] = 2
	fmt.Printf("%p", res)
	return res
}
