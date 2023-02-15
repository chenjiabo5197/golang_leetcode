package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(missingNumber(nums))
}

// 剑指 Offer 53 - II. 0～n-1中缺失的数字
func missingNumber(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		// 索引所在的值和索引相等则证明前半部分不缺数，后半部分缺数
		if nums[mid] == mid {
			left = mid + 1
		} else {
			// 其余情况都是前半部分缺数
			right = mid - 1
		}
	}
	return left
}
