package main

import (
	"fmt"
	"github.com/philchia/agollo"
)

func main() {
	err := agollo.StartWithConfFile("app.properties")
	if err != nil {
		fmt.Println(err)
		return
	}
	// val := agollo.GetStringValue("consul.client", "unknow")
	// fmt.Println(val)

	// events := agollo.WatchUpdate()
	// changeEvent := <-events
	// bytes, _ := json.Marshal(changeEvent)
	// fmt.Println("event:", string(bytes))

	// 非kv形式的配置
	res := agollo.GetNameSpaceContent("application", "unknow")
	fmt.Println(res)

	res2 := agollo.GetAllKeys("application")
	fmt.Println(res2)

	// 没这个方法了？
	// agollo.SubscribeToNamespaces("newNamespace1", "newNamespace2")
}
