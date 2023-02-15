package main

import "fmt"

func main()  {
	nums := []int{1,3,5,6}
	target := 0
	fmt.Println(searchInsert(nums, target))
}

//35. 搜索插入位置  O(log n) 的算法 二分查找
func searchInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, low, high, target int) int {
	if low > high {
		return high+1
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