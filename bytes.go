package main

import "fmt"

func main() {
    s := "abc"
    s1 := []byte(s)
    s1[0] = 1
    s2 := []byte(s)

    s3 := [...]int{1,2,3,4}
    s4 := s3[:]
    s4[0] = 0
    s5 := s3[:]
    fmt.Printf("%p %p %p", &s, &s1, &s2)
    fmt.Printf("%v %v %v", s3, s4, s5)
}
