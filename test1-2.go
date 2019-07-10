package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name string
	Age  int
}

func (b *Bird) Fly()  {
	fmt.Println("flying....")
}

func main() {
	sparrow := &Bird{"Sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()

	typeOft := s.Type()
	for i :=0; i<s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOft.Field(i).Name, f.Type(), f.Interface())
	}
}
