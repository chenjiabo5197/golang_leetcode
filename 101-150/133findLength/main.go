package main

import "fmt"

func main() {
	nums1 := []int{0, 1, 1, 1, 1}
	nums2 := []int{1, 0, 1, 0, 1}
	fmt.Println(findLength(nums1, nums2))
}

//718. 最长重复子数组
func findLength(nums1 []int, nums2 []int) int {
	arr := make([][]int, len(nums1))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, len(nums2))
	}
	if nums1[0] == nums2[0] {
		arr[0][0] = 1
	} else {
		arr[0][0] = 0
	}
	maxLength := 0
	for i := 1; i < len(arr); i++ {
		if nums1[i] == nums2[0] {
			arr[i][0] = 1
		}
		maxLength = max(maxLength, arr[i][0])
	}
	for j := 1; j < len(arr[0]); j++ {
		if nums1[0] == nums2[j] {
			arr[0][j] = 1
		}
		maxLength = max(maxLength, arr[0][j])
	}
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[0]); j++ {
			if nums1[i] == nums2[j] {
				arr[i][j] = arr[i-1][j-1] + 1
			} else {
				arr[i][j] = 0
			}
			maxLength = max(maxLength, arr[i][j])
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
