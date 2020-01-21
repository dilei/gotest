package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
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

	var nums = []int{4, 4, 4, 1, 4}
	fmt.Println(subsets(nums))

}
func subsets(nums []int) [][]int {
	sort.Ints(nums)
	mm := make(map[string]bool)
	res := [][]int{}
	tmp := []int{}
	for i := 0; i <= len(nums); i++ {
		dfs(nums, &res, tmp, i, 0, mm)
	}
	return res
}

func slice2string(nums []int) string {
	strs := make([]string, len(nums))
	for _, v := range nums {
		strs = append(strs, strconv.Itoa(v))
	}
	return strings.Join(strs, "")
}
func dfs(nums []int, res *[][]int, tmp []int, k, start int, mm map[string]bool) {
	if len(tmp) == k {
		str := slice2string(tmp)
		if _, ok := mm[str]; ok {
			// 去重
		} else {
			cur := make([]int, k)
			copy(cur, tmp)
			*res = append(*res, cur)
			mm[str] = true
		}
		return
	}
	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		dfs(nums, res, tmp, k, i+1, mm)
		tmp = tmp[:len(tmp)-1]
	}
}
