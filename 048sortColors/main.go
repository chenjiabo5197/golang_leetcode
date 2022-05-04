package main

import "fmt"

/*
颜色分类
 */

func main() {
	data := []int{1, 2, 0}
	sortColors(data)
}

//快排的思想
func sortColors(nums []int)  {
	index := 0
	left := 0
	right := len(nums)-1
	for index<=right {
		if nums[index] == 2 {
			nums[index], nums[right] = nums[right], nums[index]
			right--
		}else if nums[index] == 0 {
			nums[index], nums[left] = nums[left], nums[index]
			index++
			left++
		}else {
			index++
		}
	}
	fmt.Println(nums)
}