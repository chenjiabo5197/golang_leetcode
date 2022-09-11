package main

import "fmt"

func main() {
	arr := []int{1,1,2}
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	tempMap := make(map[int]int)
	for i:=0;i<len(nums);i++ {
		 if _, ok := tempMap[nums[i]]; !ok{
		 	tempMap[nums[i]] = nums[i]
		 }else {
		 	nums = append(nums[:i], nums[i+1:]...)
		 	i--
		 }
	}
	return len(nums)
}
