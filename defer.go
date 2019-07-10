package main

/*
return xxx这一条语句并不是一条原子指令!
函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。
defer表达式可能会在设置函数返回值之后，在返回到调用函数之前，修改返回值，使最终的函数返回值与你想象的不一致。
其实使用defer时，用一个简单的转换规则改写一下，就不会迷糊了。改写规则是将return语句拆成两句写，return xxx会被改写成:
返回值 = xxx
调用defer函数
空的return
*/

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
	println(add(1, 2))  // 输出: 103
	println(add2(1, 2)) // 输出: 3
}
