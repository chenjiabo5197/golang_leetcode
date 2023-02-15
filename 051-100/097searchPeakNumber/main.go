package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(nums))
}

// 162. 寻找峰值  时间复杂度log n
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		// 结果和(left+right)/2 结果相同，但是left和right同时过大有可能会导致溢出
		mid := left + (right-left)/2
		// 寻找一个处于爬坡状态的值，由于nums[-1]和num[len(nums)]的值均为 负无穷，所以处于爬坡状态的值一定可以找到一个峰值
		if nums[mid] < nums[mid+1] { //爬坡状态，往后面继续找
			left = mid + 1
		} else { // 非爬坡状态，往前找，非爬坡状态不-1是因为有可能当前数即为峰值，而爬坡状态已经确认过当前数非峰值
			right = mid
		}
	}
	// 此时left和right相等
	return left
}
