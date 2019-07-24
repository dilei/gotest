package main

import "fmt"

func main() {
	var data interface{}

	data = []string{"1", "2", "3"}

	switch data.(type) {
	case int:
		fmt.Println("int")
	case []string:
		fmt.Println("string slice")
	}
}
