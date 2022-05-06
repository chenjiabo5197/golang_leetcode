package main

import (
	"fmt"
)

/*
有序数组的平方
 */

func main() {
	nums :=[]int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
}

// 暴力解法，不美观
//func sortedSquares(nums []int) []int {
//	result := make([]int, len(nums))
//	for i:=0;i<len(nums);i++ {
//		result[i] = nums[i] * nums[i]
//	}
//	sort.Ints(result)
//	return result
//}

// 两个指针，从头和尾遍历一次即可
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	start := 0
	end := len(nums)-1
	index := len(nums)-1
	for start<=end {
		startNum := nums[start] * nums[start]
		endNum := nums[end] * nums[end]
		if startNum > endNum {
			result[index] = startNum
			start++
		}else {
			result[index] = endNum
			end--
		}
		index--
	}
	return result
}