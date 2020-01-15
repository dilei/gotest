package main

import "fmt"

func main() {
	s := make([]int, 2)
	// mdSlice(s)  // [0 0]
	mdSlice2(&s) // [0 0 1 2]
	fmt.Println(s)
}

func mdSlice(s []int) {
	s = append(s, 1)
	s = append(s, 2)
}

func mdSlice2(s *[]int) {
	*s = append(*s, 1)
	*s = append(*s, 2)
}
