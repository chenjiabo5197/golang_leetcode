package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

var (
	res   [][]int
	path  []int
	isUse []bool
)

//46. 全排列
func permute(nums []int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	isUse = make([]bool, len(nums))
	for i := 0; i < len(isUse); i++ {
		isUse[i] = false
	}
	solve(nums)
	return res
}

func solve(nums []int) {
	if isErgodic() {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
	} else {
		for i := 0; i < len(nums); i++ {
			if !isUse[i] {
				isUse[i] = true
				path = append(path, nums[i])
				solve(nums)
				isUse[i] = false
				path = path[:len(path)-1]
			}
		}
	}
}

// 判断所有数字是否已经遍历完，完成返回true，否则返回false
func isErgodic() bool {
	for i := 0; i < len(isUse); i++ {
		if !isUse[i] {
			return false
		}
	}
	return true
}
