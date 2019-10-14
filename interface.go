package main

import (
	"fmt"
)

func main() {
	var data interface{}

	data = []string{"1", "2", "3"}
	data = nil
	data = []int64{1, 2, 3}

	data = "123"

	// 如果类型断言错误 typeErr为int64的零值
	// 下面的方式会出现panic
	// typeErr := data.(int)
	// fmt.Println(typeErr)

	// 正确方式
	// typeErr, ok := data.(int64)
	// if ok {
	// 	fmt.Println(typeErr)
	// }

	switch data.(type) {
	case int:
		fmt.Println("int")
	case []int64:
		fmt.Println("int64 slice")
	case []string:
		fmt.Println("string slice")
	case nil:
		fmt.Println("nil")
	}

	a := Newdog()
	a.say()

	i := make(map[string]interface{})
	i["abc"] = nil
	delete(i, "ddd")
	if _, ok := i["abc"]; ok {
		fmt.Println("nil is ok")
	}

	var strArr []int
	// strArr = nil
	// fmt.Println(len(strArr))

	strArr = []int{1, 2, 3}
	// fmt.Println(append(strArr[:1], strArr[2:]...))
	// fmt.Println(strArr[:0], strArr[1:])
	// fmt.Println(append(strArr[:2], strArr[2:]...))

	fmt.Println(len(strArr), cap(strArr))
	for key := range strArr {
		if key == 1 {
			strArr = append(strArr[:3], strArr[3:]...)
		}
	}
	fmt.Println(strArr)
	fmt.Println(strArr[len(strArr):])

}

type animal interface {
	say()
	run()
}

type dog struct {
	Age int
}

func Newdog() animal {
	return &dog{}
}

func (*dog) say() {
	fmt.Println("wang wang wang")
}
func (*dog) run() {
	fmt.Println("runing")
}
