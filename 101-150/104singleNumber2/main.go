package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(nums))
}

//137. 只出现一次的数字 II
func singleNumber(nums []int) int {
	isCompare := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		value, ok := isCompare[nums[i]]
		if ok {
			isCompare[nums[i]] = value + 1
		} else {
			isCompare[nums[i]] = 1
		}
	}
	for key, value := range isCompare {
		if value != 3 {
			return key
		}
	}
	return 0
}

//func singleNumber(nums []int) int {
//	isCompare := make(map[int]int)
//	for i:=0;i<len(nums);i++ {
//		if _, ok := isCompare[nums[i]]; ok {
//			continue
//		}
//		for j:=i+1;j<len(nums);j++ {
//			if nums[i] == nums[j] {
//				isCompare[nums[i]] = nums[i]
//				break
//			}
//			if nums[i] != nums[j] && j == len(nums) -1 {
//				return nums[i]
//			}
//		}
//	}
//	return nums[len(nums) - 1]
//}
