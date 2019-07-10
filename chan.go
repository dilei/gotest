package main

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
}
