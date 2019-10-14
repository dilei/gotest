package main

import "fmt"

func main() {
	ps := PersonS{AdminS{Name: "dilei"}, UserS{"dilei2"}, 23, struct{}{}}
	// Ambiguous reference 'Name'
	// fmt.Println(ps.Name)

	fmt.Println(ps.UserS.Name, ps.AdminS.Name)

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
	AdminS
	UserS
	Age int
	Obj struct{}
}

type AdminS struct {
	Name   string
	Member map[int64]int64
}

type UserS struct {
	Name string
}
