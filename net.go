package main

import (
	"net"
	"time"
)

func main() {
	net.DialTimeout("tcp", "mynewproddb.idc2:3306", time.Millisecond*10)

}
