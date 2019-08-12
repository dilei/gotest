package main

import (
	"fmt"
	"regexp"
)

func main() {

	// text := `满100-50##满40-20`
	text := `满##满`

	reg := regexp.MustCompile(`[0-9-]+$`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
}
