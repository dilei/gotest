package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	fmt.Println(file)
	fmt.Println(".." + string(filepath.Separator))
	fmt.Println(filepath.Join(file, ".."+string(filepath.Separator)))
	fmt.Println(apppath)

	fmt.Println(filepath.Abs(filepath.Dir(os.Args[0])))
	fmt.Println(os.Getwd())
}
