package main

import (
	"fmt"
	"io"
	"net"
)

const RECV_BUF_LEN = 10

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:5566")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("server starting")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(conn.RemoteAddr())

		go EchoServer(conn)
	}
}

func EchoServer(conn net.Conn) {
	buf := make([]byte, RECV_BUF_LEN)
	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		switch err {
		case nil:
			conn.Write(buf[0:n])
		case io.EOF:
			fmt.Println("End of data %s", err)
		default:
			fmt.Println("Error: %s", err)
		}
	}
}
