package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(nums, target))
}

var (
	res [][]int
)

//40. 组合总和 II
func combinationSum2(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(candidates)
	DFS(candidates, path, target, 0)
	return res
}

// DFS 深度优先
func DFS(candidates, path []int, sum, index int) {
	if sum == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > sum {
			break
		}
		if i != index && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		DFS(candidates, path, sum-candidates[i], i+1)
		path = path[:len(path)-1]
	}
}
