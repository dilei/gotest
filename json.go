package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Server struct {
	ServerName string
	ServerIP   string
	Id         int `json:"id,string"`
	Num        int `json:"num"`
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	// ok
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1","id":"1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2","id":"2"}]}`

	// json: invalid use of ,string struct tag, trying to unmarshal unquoted value into int
	// str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1","id":1},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2","id":2}]}`

	// json: cannot unmarshal string into Go struct field Server.Servers.num of type int
	// str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1","id":"1","num":"123"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2","id":"2","num":"1234"}]}`

	// json: cannot unmarshal number into Go struct field Server.Servers.ServerIP of type string
	// str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":12312,"id":"1","num":"123"},{"serverName":"Beijing_VPN","serverIP":1231,"id":"2","num":"1234"}]}`

	fmt.Println(s)
	var ss interface{}
	str = `{}`
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(s)

	switch ss.(type) {
	case Serverslice:
		fmt.Println("serverslice")
	default:
		fmt.Println("not serverslice")
	}

	byteArr, _ := json.Marshal(s)
	fmt.Println(string(byteArr))

	fmt.Println(fmt.Sprintf("%.02f", 1.2))
	fmt.Println(fmt.Sprintf("%.2f", 1.2))
	val, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 1.3), 64)
	fmt.Println(val)

}
