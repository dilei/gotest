package main

import "fmt"

func main() {
	res := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(res)
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	cur := []int{}
	dfs2(candidates, target, 0, &cur, &res)
	return res
}

func dfs2(candidates []int, target int, start int, cur *[]int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
	}
	if target < 0 {
		return
	}
	for i := start; i < len(candidates); i++ {
		*cur = append(*cur, candidates[i])
		dfs2(candidates, target-candidates[i], i, cur, res)
		*cur = (*cur)[:len(*cur)-1]
	}
}
