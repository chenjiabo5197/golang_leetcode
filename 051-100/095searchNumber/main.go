package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 1}
	target := 8
	fmt.Println(search(nums, target))
}

// 剑指 Offer 53 - I. 在排序数组中查找数字 I
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	index := binarySearch(nums, target, left, right)
	leftIndex, rightIndex := 0, 0
	if index != -1 {
		leftIndex, rightIndex = index, index
	}
	for index != -1 {
		tempLeftIndex := binarySearch(nums, target, left, leftIndex-1)
		tempRightIndex := binarySearch(nums, target, rightIndex+1, right)
		if tempLeftIndex == -1 && tempRightIndex == -1 {
			index = -1
		}
		if tempLeftIndex != -1 {
			leftIndex = tempLeftIndex
		}
		if tempRightIndex != -1 {
			rightIndex = tempRightIndex
		}
	}
	if rightIndex == 0 && nums[0] != target {
		return 0
	}
	return rightIndex - leftIndex + 1
}

func binarySearch(nums []int, target, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
