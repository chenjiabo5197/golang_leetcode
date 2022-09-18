package main

import (
	"fmt"
)

func main()  {
	nums := []int{4,5,6,7,0,1,2}
	target := 0
	fmt.Println(search(nums, target))
}

// 33. 搜索旋转排序数组 旋转之后的数组从middle分开，一半有序，一半无序
func search(nums []int, target int) int {
	low := 0
	high := len(nums)-1
	for low <= high {
		middle := (low + high) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[low] <= nums[middle] { // 左侧区间有序
			if nums[low] <= target && nums[middle] >= target {  // 目标在左侧区间
				high = middle - 1
			}else {
				low = middle + 1
			}
		}else { // 右侧区间有序
			if nums[middle] <= target && nums[high] >= target {   // 目标在右侧区间
				low = middle + 1
			}else {
				high = middle - 1
			}
		}
	}
	return -1
}

