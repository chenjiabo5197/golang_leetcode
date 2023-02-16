package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

//53. 最大子数组和
func maxSubArray(nums []int) int {
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > 0 {
			// arr[i] 表示以第i位结尾的最大子数组之和
			arr[i] = arr[i-1] + nums[i]
		} else {
			arr[i] = nums[i]
		}
	}
	max := nums[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
