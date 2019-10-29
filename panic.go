package main

import (
    "fmt"
    "golang.org/x/net/html"
)

func main() {

}

// 如果没有标题则返回错误
func soleTitle(doc *html.Node) (title string, err error) {
    type bailout struct{}
    defer func() {
        switch p := recover(); p{
        case nil:
            // 没有56 机
        case bailout{}:
            // ” 预期的” 窘机
            err = fmt.Errorf("multiple title eleme nts")
        default:
            panic(p) //未预期的后机；继续岩机过程
        }
    }()
    if title != "" {
        panic(bailout{}) //多个标题元素
    }
    return title, err
}
