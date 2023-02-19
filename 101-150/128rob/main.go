package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

//198. 打家劫舍
/*
	一维数组arr，arr[i]表示到i为止，可以获取大最大金额
	arr[i] = max(arr[i-2]+nums[i], arr[i-1])
*/
func rob(nums []int) int {
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	for i := 1; i < len(arr); i++ {
		if i == 1 {
			arr[i] = max(nums[0], nums[1])
		} else if i >= 2 {
			arr[i] = max(arr[i-2]+nums[i], arr[i-1])
		}
	}
	return arr[len(arr)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
