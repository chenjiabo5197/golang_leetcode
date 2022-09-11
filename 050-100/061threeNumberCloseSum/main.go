package main

import (
	"fmt"
	"math"
	"sort"
)

/*
最接近的三数之和
 */

func main() {
	nums := []int{-1,2,1,-4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

//解法与三数之和差不多，遍历完寻找一个最接近的和
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	curIndex := 0
	minDis := float64(65535)
	minSum := 0
	for curIndex < len(nums)-1 {
		leftIndex := curIndex+1
		rightIndex := len(nums)-1
		for leftIndex < rightIndex {
			sum := nums[curIndex] + nums[leftIndex] + nums[rightIndex]
			dis := math.Abs(float64(sum - target))
			if dis < minDis {
				minSum = sum
				minDis = dis
			}
			if sum > target {
				rightIndex--
			}else if sum < target {
				leftIndex++
			}else {
				return target
			}
		}
		curIndex++
	}
	return minSum
}