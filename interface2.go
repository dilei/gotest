package main

import (
	"fmt"
	"reflect"
)

type State struct{}

func testnil1(a, b interface{}) bool {
	fmt.Println(a)
	return a == b
}

func testnil2(a *State, b interface{}) bool {
	fmt.Println(a)
	return a == b
}

func testnil3(a interface{}) bool {
	fmt.Println(a)
	return a == nil
}

func testnil4(a *State) bool {
	fmt.Println(a)
	return a == nil
}

func testnil5(a interface{}) (bool, bool) {
	fmt.Println(a)
	v := reflect.ValueOf(a)
	return !v.IsValid(), v.IsNil()
}

func main() {
	var a *State
	fmt.Println(testnil1(a, nil))
	fmt.Println(testnil2(a, nil))
	fmt.Println(testnil3(a))
	fmt.Println(testnil4(a))
	fmt.Println(testnil5(a))
}
