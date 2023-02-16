package main

import "fmt"

func main() {
	k := 3
	n := 7
	fmt.Println(combinationSum3(k, n))
}

var (
	res  [][]int
	path []int
)

//216. 组合总和 III
func combinationSum3(k int, n int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	solve(n, k, 10)
	return res
}

func solve(sum, nums, index int) {
	if sum == 0 && nums == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
	} else {
		for i := index - 1; i >= 1; i-- {
			if i > sum {
				continue
			}
			path = append(path, i)
			solve(sum-i, nums-1, i)
			path = path[:len(path)-1]
		}
	}
}
