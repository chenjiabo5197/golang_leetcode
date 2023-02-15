package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}

//268. 丢失的数字
func missingNumber(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	originSum := 0
	if len(nums)%2 == 0 { //从0开始的数字数量是偶数个
		originSum = len(nums) / 2 * (len(nums) + 1)
	} else { //数字数量是奇数个
		originSum = len(nums) * (len(nums)/2 + 1)
	}
	return originSum - sum
}
