package main

import (
	"fmt"
	"math"
)

/*
盛水最多的容器
 */

func main() {
	data := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(data))
}

// 暴力解法  超时
//func maxArea(height []int) int {
//	max := -1
//	for i:=0;i<len(height);i++ {
//		for j:=0;j<len(height);j++ {
//			area := int(math.Min(float64(height[i]), float64(height[j])) * float64(j - i))
//			if area > max {
//				max = area
//			}
//		}
//	}
//	return max
//}

// 用两个指针分别指向头和尾，移动指向线段较短的指针，使宽度变小的情况下看高度是否会变高
func maxArea(height []int) int {
	start := 0
	end := len(height) -1
	max := -1
	for start < end {
		area := (end-start) * int(math.Min(float64(height[start]), float64(height[end])))
		if height[start] < height[end] {
			start++
		}else {
			end--
		}
		if max < area {
			max = area
		}
	}
	return max
}