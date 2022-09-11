package main

import "fmt"

func main() {
	num := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(num))
}


func findRepeatNumber(nums []int) int {
	temp := make(map[int]bool)
	for i:=0;i<len(nums);i++ {
		if _,ok := temp[nums[i]]; !ok {
			temp[nums[i]] = true
		}else {
			return nums[i]
		}
	}
	return 0
}