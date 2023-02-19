package main

import "fmt"

func main() {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums))
}

//300. 最长递增子序列
/*
	一维数组arr，arr[i]表示以i结尾时，最长的递增子序列
*/
func lengthOfLIS(nums []int) int {
	arr := make([]int, len(nums))
	for i := 0; i < len(arr); i++ {
		arr[i] = 1
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				arr[i] = max(arr[i], arr[j]+1)
			}
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
