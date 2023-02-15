package main

import (
	"fmt"
	"sort"
)

/*
三数之和
 */

func main() {
	data := []int{-2, 0, 0, 2, 2}
	fmt.Println(threeSum(data))
}

// 先将数组排序，然后从小到大遍历数组，设置3个索引，curIndex，leftIndex，rightIndex，leftIndex初始位置位于curIndex的右邻边，rightIndex初始位置
// 位于数组的最右边，然后开始遍历数组，计算3个索引指向的数字之和，根据和决定leftIndex和rightIndex的移动
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	curIndex := 0
	sort.Ints(nums)
	for curIndex<len(nums)-1 {
		if nums[curIndex] > 0 {   //当前索引数大于0，则三数之和肯定大于0，循环结束
			break
		}
		leftIndex := curIndex+1
		rightIndex := len(nums)-1
		for leftIndex < rightIndex {
			if curIndex != 0 && nums[curIndex] == nums[curIndex-1] {  //执行的是去重的操作
				break
			}
			if nums[curIndex] + nums[leftIndex] + nums[rightIndex] > 0 {
				rightIndex--
			} else if nums[curIndex] + nums[leftIndex] + nums[rightIndex] < 0 {
				leftIndex++
			} else {
				tempSlice := []int{nums[curIndex], nums[leftIndex], nums[rightIndex]}
				leftIndex++
				for leftIndex < rightIndex && nums[leftIndex] == nums[leftIndex-1] {  //执行的是去重的操作
					leftIndex++
				}
				result = append(result, tempSlice)
			}
		}
		curIndex++
	}
	return result
}