package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))
}

//260 只出现一次的数字 III
func singleNumber(nums []int) []int {
	isCompare := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		value, ok := isCompare[nums[i]]
		if ok {
			isCompare[nums[i]] = value + 1
		} else {
			isCompare[nums[i]] = 1
		}
	}
	result := make([]int, 0)
	for key, value := range isCompare {
		if value != 2 {
			result = append(result, key)
		}
	}
	return result
}
