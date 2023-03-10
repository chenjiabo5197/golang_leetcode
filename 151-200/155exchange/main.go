package main

import "fmt"

func main() {
	nums := []int{1, 3, 5}
	fmt.Println(exchange(nums))
}

//剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
/*
	两个下标，一个从前开始遍历，一个从后开始遍历
*/
func exchange(nums []int) []int {
	start := 0
	end := len(nums) - 1
	for start < end {
		for start < end && nums[start]%2 != 0 {
			start++
		}
		for start < end && nums[end]%2 == 0 {
			end--
		}
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
	return nums
}
