package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	ps := PersonS{UserS{"dilei2", "1231"}, "ddd", 23, struct{}{}}
	// Ambiguous reference 'Name'
	// fmt.Println(ps.Name)

	byteArr, err := json.Marshal(ps)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(byteArr))

	ps2 := PersonS2{UserS{"dilei2", "123"}, 23}
	byteArr, err = json.Marshal(ps2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(byteArr))

	as := AdminS{}
	fmt.Println(as.Member == nil)
	for key, val := range []int64{1, 2, 3} {
		if as.Member == nil {
			mem := make(map[int64]int64)
			mem[12] = 123
			as.Member = mem
		} else {
			as.Member[int64(key)] = val
		}
	}

	fmt.Println(as.Member)
	as.Name = "123"
	fmt.Println(as.Name)

	setName(as)
	fmt.Println(as.Name)

	age := 2
	isOne := age == 1
	fmt.Println(isOne)

	fmt.Println(ps)

	obj := struct{}{} // 空结构体不占用内存空间
	fmt.Println(obj)

}

func setName(as AdminS) {
	as.Name = "456"
}

type PersonS struct {
	UserS
	Name string
	Age  int
	Obj  struct{}
}

type PersonS2 struct {
	UserS
	Age int
}
type AdminS struct {
	Name   string
	Member map[int64]int64
}

type UserS struct {
	Name  string
	Email string
}

func (u UserS) GetName() string {
	return "name"
}

type U2 UserS

var u2 U2
