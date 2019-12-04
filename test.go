package main

import (
	"encoding/json"
	"fmt"
)

type childProd map[string]interface{}

func main() {
	tmp := make(map[int64]map[string]interface{})
	tmp[123] = map[string]interface{}{
		"aa": 123,
	}
	tmp[321] = map[string]interface{}{
		"aa": 123,
	}
	// tmp[123]["ebook"] = tmp[123]
	child1 := make(map[string]interface{})
	for k, v := range tmp[123] {
		child1[k] = v
	}

	child2 := make(map[string]interface{})
	for k, v := range tmp[321] {
		child2[k] = v
	}

	tmp[123]["ebook"] = child1
	// tmp[321]["book"] = tmp[123]
	tmp[321]["book"] = child2
	byteArr, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(byteArr))

}
