package main

import "fmt"

const (
	a1 = iota
	b2
	c3
	d4 = 'c'
	e5 = iota
)

func main() {
	fmt.Println(a1, b2, c3, d4, e5)
}
