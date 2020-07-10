package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "abc"
	s1 := []byte(s)
	s1[0] = 1
	s2 := []byte(s)

	s3 := [...]int{1, 2, 3, 4}
	s4 := s3[:]
	s4[0] = 0
	s5 := s3[:]
	fmt.Printf("%p %p %p", &s, &s1, &s2)
	fmt.Println()
	fmt.Printf("%v %v %v", s3, s4, s5)
	fmt.Println()

	a := 157
	b := byte(a)
	var c rune = 'c'
	var d byte = 'c'
	e := 'c'
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))
	// fmt.Println(c == d)  // mismatched types
	// fmt.Println(e == d) // mismatched types
	fmt.Println(e == c)
	fmt.Println(b)
}
