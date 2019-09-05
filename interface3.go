package main

import "fmt"

type update interface {
	setId(i int)
	getId() int
}

type stu struct {
	id int
}

func (s stu) setId(i int) {
	fmt.Printf("%p", &s)
	fmt.Printf("%s", "---")
	s.id = i
}

func (s stu) getId() int {
	fmt.Printf("%p", &s)
	fmt.Printf("%s", "---")
	return s.id
}

type stu2 struct {
	id int
}

func (s *stu2) setId(i int) {
	fmt.Printf("%p", s)
	fmt.Printf("%s", "---")
	s.id = i
}

func (s *stu2) getId() int {
	fmt.Printf("%p", s)
	fmt.Printf("%s", "---")
	return s.id
}

func NewStu() update {
	return &stu{}
}
func NewStu2() update {
	return &stu2{}
}

func main() {
	s := NewStu()
	fmt.Printf("%p", &s)
	fmt.Printf("%s", "---")
	s.setId(10)
	fmt.Printf("%p", &s)
	fmt.Printf("%s", "---")
	fmt.Println(s.getId())

	s2 := NewStu2()
	fmt.Printf("%p", s2)
	fmt.Printf("%s", "---")
	s2.setId(10)
	fmt.Printf("%p", s2)
	fmt.Printf("%s", "---")
	fmt.Println(s2.getId())
}
