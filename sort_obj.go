package main

import (
	"fmt"
	"sort"
	"strings"
)

type AttribNameValue struct {
	Name string
	Age  int
	email  string
}

type AttribNameValueArray []AttribNameValue

func (p AttribNameValueArray) Len() int {
	return len(p)
}

func (p AttribNameValueArray) Less(i, j int) bool {
	if p[i].Name < p[j].Name {
		return true
	}
	if p[i].Name > p[j].Name {
		return false
	}

	return p[i].Age < p[j].Age
}

func (p AttribNameValueArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *AttribNameValueArray) Sort() {
	sort.Sort(p)
}

func main() {
	n1 := AttribNameValue{"dilei", 9, "111"}
	n2 := AttribNameValue{"dilei", 9, "222"}
	n3 := AttribNameValue{"dilei", 9, "333"}
	n4 := AttribNameValue{"dilei", 9, "444"}

	nArr := AttribNameValueArray{n1, n2, n3, n4}

	fmt.Println(nArr)
	nArr.Sort()
	fmt.Println(nArr)

	strArr := strings.Split("", ",")
	fmt.Println(strArr[0])

	mapObj := make(map[int]map[int]interface{})
	tmp := make(map[int]interface{})
	tmp[1] = 123
	mapObj[1] = tmp

	var res int
	res = mapObj[0][123].(int)
	fmt.Println(res)
}

