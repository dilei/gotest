package main

import (
	"fmt"
	"log"
	"time"
)

/*
return xxx这一条语句并不是一条原子指令!
函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。
defer表达式可能会在设置函数返回值之后，在返回到调用函数之前，修改返回值，使最终的函数返回值与你想象的不一致。
其实使用defer时，用一个简单的转换规则改写一下，就不会迷糊了。改写规则是将return语句拆成两句写，return xxx会被改写成:
返回值 = xxx
调用defer函数
空的return
*/

func f12() (result int) {
	defer func() {
		result++
	}()
	// 先把10 赋值给result，然后再++， 结果是11
	return 10
}

/*
   z   // 无需创建新的返回变量
   z += 100
   return
*/
func add(x, y int) (z int) {
	defer func() {
		z += 100
	}()
	z = x + y
	return
}

/*
	r := z  // 需要先创建新的返回变量
	z += 100
	return
*/
func add2(x, y int) int {
	var z int
	defer func() {
		z += 100
	}()
	z = x + y
	return z
}

func main() {
	println(f12())  // 输出:
	println(add(1, 2))  // 输出: 103
	println(add2(1, 2)) // 输出: 3

	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // 别忘记这对圆括号
	// defer trace("bigSlowOperation") // 没有圆括号
	defer fmt.Print(132)
	// 这里是一些处理
	time.Sleep(10 * time.Second)  // 通过休眠仿真慢操作
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func (){log.Printf("exit %s (%s)", msg, time.Since(start))}
}
