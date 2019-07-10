package main

import (
	"fmt"
	"net"
	"time"
)

const POST_BUF_LEN = 1024

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5566")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	buf := make([]byte, POST_BUF_LEN)

	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Hello %03d", i)
		n, err := conn.Write([]byte(msg))
		if err != nil {
			println("write error", err.Error())
			break
		}
		fmt.Println(msg)

		n, err = conn.Read(buf)
		if err != nil {
			println("Read error", err.Error())
			break
		}
		fmt.Println(buf[0:n])

		time.Sleep(time.Second)
	}
}
