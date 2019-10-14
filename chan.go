package main

import (
	"fmt"
)

/*
func main() {
// 示例1。
ch1 := make(chan int, 1)
ch1 <- 2
<- ch1

// 示例2。
ch2 := make(chan int, 1)
//elem, ok := <-ch2 // 通道已空，因此这里会造成阻塞。
//_, _ = elem, ok
ch2 <- 1

// 示例3。
var ch3 chan int
//ch3 <- 1 // 通道的值为nil，因此这里会造成永久的阻塞！
//<-ch3 // 通道的值为nil，因此这里会造成永久的阻塞！
_ = ch3

}
*/

func main() {
	// deadlock
	// 对于非缓冲通道，情况要简单一些。无论是发送操作还是接收操作，一开始执行就会被阻塞，
	// 直到配对的操作也开始执行，才会继续传递。由此可见，非缓冲通道是在用同步的方式传递数据。
	// 也就是说，只有收发双方对接上了，数据才会被传递。
	// 并且，数据是直接从发送方复制到接收方的，中间并不会用非缓冲通道做中转。相比之下，缓冲通道则在用异步的方式传递数据。
	// c := make(chan int)
	// c <- 1
	// <-c

	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("first element: %v", elem1)
	fmt.Println(fmt.Sprintf("%.02f", 12.1))

	var ch chan bool
	ch <- true
	go func() {
		defer func() {
			if x := recover(); x != nil {
				fmt.Println("error")
			}
		}()
		panic("123")
		ch <- true
		return
	}()
	i, ok := <-ch
	if ok {
		fmt.Println(true)
	}
	fmt.Println(i)

	/*
		c := make(chan int)
		o := make(chan bool)

		// <- o // 1、如果直接读空通道，会直接报deadlock错误

		go func() {
			for {
				select {
				case v := <- c:
					println(v)
				case <- time.After(5 * time.Second):
					println("timeout")
					o <- true  // 相当于一个标记，通道可读，函数执行结束
					break
				}
			}
		}()
		<- o  // 2、当上面有一个并发线程时，他的作用是阻塞函数直到通道可读，即使线程没有写通道
	*/

	/*
		ch2 := make(chan int, 2)
		// close(ch2)  // 关闭通道会导致在获取时返回零值，相应使用select时，每次都会选中
		ch2 <- 1

		var intV int
		select {
		case intV = <-ch2:
			// if ok {
			// 	fmt.Println("ch2 selected")
			// } else {
			// 	fmt.Println("ch2 closed")
			// }
			fmt.Println("ch2 selected")
		case <-time.After(2 * time.Second): // 2s超时
		}
		fmt.Println(intV)
	*/
}
