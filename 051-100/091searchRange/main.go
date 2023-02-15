package main

import "fmt"

func main()  {
	nums := []int{2,2}
	target := 2
	fmt.Println(searchRange(nums, target))
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	firstIndex := binarySearch(nums, 0, len(nums)-1, target)
	if firstIndex == -1 {
		return []int{-1, -1}
	}else {
		lowIndex := firstIndex
		lastLowIndex := firstIndex
		for lowIndex != -1 {
			lowIndex = binarySearch(nums, 0, lowIndex-1, target)
			if lowIndex != -1 {
				lastLowIndex = lowIndex
			}
		}
		highIndex := firstIndex
		lastHighIndex := firstIndex
		for highIndex != -1 {
			highIndex = binarySearch(nums, highIndex+1, len(nums)-1, target)
			if highIndex != -1 {
				lastHighIndex = highIndex
			}
		}
		return []int{lastLowIndex, lastHighIndex}
	}
}

func binarySearch(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	middle := (low + high) / 2
	if nums[middle] == target {
		return middle
	}else if nums[middle] > target {
		return binarySearch(nums, low, middle-1, target)
	}else {
		return binarySearch(nums, middle+1, high, target)
	}
}