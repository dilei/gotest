package main

import (
    "fmt"
    "time"
)

func main() {
    for i:=0; i<10; i++{
        go func(){
            processValue(i)
        }()
    }

    for i:=0; i<10; i++{
        go func(i int){
            processValue(i)
        }(i)
    }
    time.Sleep(time.Second)
}

func processValue(i int){
    fmt.Println(i)
}
