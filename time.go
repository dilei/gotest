package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Unix())

	fmt.Println(t.Format("20060102150405"))
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format("15:04:05"))
	fmt.Println(t.Format("15:04:05.000"))

	// Ymdhis
	ymdhis := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println(ymdhis)
}
