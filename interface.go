package main

import "fmt"

func main() {
	var data interface{}

	data = []string{"1", "2", "3"}
	data = nil
	data = []int64{1, 2, 3}

	// 如果类型断言错误 typeErr为int64的零值
	// 下面的方式会出现panic
	// typeErr := data.(int64)
	typeErr, ok := data.(int64)
	if ok {
		fmt.Println(typeErr)
	}

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
