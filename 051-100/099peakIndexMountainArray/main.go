package main

import "fmt"

func main() {
	nums := []int{0, 10, 5, 2}
	fmt.Println(peakIndexInMountainArray(nums))
}

//852. 山脉数组的峰顶索引
func peakIndexInMountainArray(arr []int) int {
	left := 1
	right := len(arr) - 2 // 数组的最左侧和最右侧不是山峰
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
		// 判断mid处于上升还是下降状态
		// 处于上升状态
		if arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
			left = mid + 1
		}
		// 处于下降状态
		if arr[mid] < arr[mid-1] && arr[mid] > arr[mid+1] {
			right = mid - 1
		}
		// arr是一个山脉数组，所以只有这三种状态
	}
	return left
}
