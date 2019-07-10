package main

import "fmt"

const (
	a = iota
	b
	c
	d = 'c'
	e = iota
)

func main()  {
	fmt.Println(a, b, c, d, e)
}