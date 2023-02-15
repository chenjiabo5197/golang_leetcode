package main

import "fmt"

/*
两数之和
 */

func main() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

// 存在问题，没办法忽略自身元素
//func twoSum(nums []int, target int) []int {
//	indexTable := make(map[int]int)
//	result := make([]int, 2)
//	for i:=0;i<len(nums);i++ {
//		indexTable[nums[i]] = i
//	}
//	for i:=0;i<len(nums);i++ {
//		subNum := target - nums[i]
//		if value, ok := indexTable[subNum]; ok {
//			result[0] = i
//			result[1] = value
//			break
//		}
//	}
//	return result
//}

func twoSum(nums []int, target int) []int {
	indexTable := make(map[int]int)
	result := make([]int, 2)
	for i:=0;i<len(nums);i++ {
		subNum := target - nums[i]
		if value, ok := indexTable[subNum]; ok {
			result[0] = i
			result[1] = value
			break
		}else {
			indexTable[nums[i]] = i
		}
	}
	return result
}