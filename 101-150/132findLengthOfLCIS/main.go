package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(nums))
}

//674. 最长连续递增序列
func findLengthOfLCIS(nums []int) int {
	arr := make([]int, len(nums))
	for i := 0; i < len(arr); i++ {
		arr[i] = 1
	}
	for i := 1; i < len(arr); i++ {
		if nums[i] > nums[i-1] {
			arr[i] = arr[i-1] + 1
		} else {
			arr[i] = 1
		}
	}
	maxLength := 0
	for i := 0; i < len(arr); i++ {
		maxLength = max(maxLength, arr[i])
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
