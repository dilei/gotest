package v1

import "fmt"

func GetAge() {
	bob := Person{age: 25, name: "Bob"}
	fmt.Print(bob.age)
}
