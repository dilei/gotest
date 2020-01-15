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

	// 模拟序列化栈溢出
	tmp[123]["ebook"] = child1
	// tmp[321]["book"] = tmp[123]
	tmp[321]["book"] = child2
	byteArr, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(byteArr))

	var nums = []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
func subsets(nums []int) [][]int {
	res := [][]int{}
	tmp := []int{}
	for i := 0; i <= len(nums); i++ {
		dfs(nums, &res, tmp, i, 0)
	}
	return res
}

func dfs(nums []int, res *[][]int, tmp []int, k, start int) {
	if len(tmp) == k {
		cur := make([]int, k)
		copy(cur, tmp)
		*res = append(*res, cur)
		return
	}
	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		dfs(nums, res, tmp, k, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}
