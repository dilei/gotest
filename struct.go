package main

import "fmt"

func main() {
	ps := PersonS{AdminS{"dilei"}, UserS{"dilei2"}, 23}
	// Ambiguous reference 'Name'
	// fmt.Println(ps.Name)

	fmt.Println(ps.UserS.Name)
	fmt.Println(ps.AdminS.Name)

}

type PersonS struct {
	AdminS
	UserS
	Age int
}

type AdminS struct {
	Name string
}

type UserS struct {
	Name string
}
