package main

import "fmt"

/*
摆动序列
 */

func main() {
	nums := []int{1,7,4,9,2,5}
	fmt.Println(wiggleMaxLength(nums))
}

// 自己想法：遍历列表，用最常增长子序列动态规划去做
//func wiggleMaxLength(nums []int) int {
//	length := make([]int, len(nums))
//	length[0] = 1
//	for i:=0;i<len(nums);i++ {
//		if i==2 {
//			if nums[i] != nums[i-1] {
//				length[i] = 2
//			}else {
//				length[i] = 1
//			}
//		}
//
//	}
//}

// 贪心思想，一个元素，3个状态，初始，上升，下降，当处于上升期时，找到最后一个上升元素，当处于下降期，找到最后一个下降元素
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	//定义3个状态
	initState := 0
	upState := 1
	downState := 2

	//开始遍历前，处于初始状态
	state := initState
	length := 1  //初始长度为1
	for i:=1;i<len(nums);i++ {
		if state == initState {  //处于初始状态
			if nums[i] > nums[i-1] {  //相等 状态不变化
				state = upState
				length++
			}else if nums[i] < nums[i-1] {
				state = downState
				length++
			}
		} else if state == upState {  //处于上升状态，只有破坏上升状态长度才会变化，一直上升状态不变，下降同理
			if nums[i] < nums[i-1] {
				state = downState
				length++
			}
		} else if state == downState {  // 处于下降状态
			if nums[i] > nums[i-1] {
				state = upState
				length++
			}
		}
	}
	return length
}