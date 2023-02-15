package main

import (
	"fmt"
)

func main()  {
	nums := []int{-1,0,3,5,9,12}
	target := 2
	fmt.Println(search(nums, target))
}

// 704. 二分查找
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	middleIndex := (low + high) /2
	if nums[middleIndex] == target {
		return middleIndex
	}else if nums[middleIndex] > target{
		return binarySearch(nums, low, middleIndex-1, target)
	}else {
		return binarySearch(nums, middleIndex+1, high, target)
	}
}
