package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2, 2, 2}
	fmt.Println(makesquare(nums))
}

//473. 火柴拼正方形
func makesquare(matchsticks []int) bool {
	sort.Ints(matchsticks)
	sum := 0
	for i := len(matchsticks) - 1; i >= 0; i-- {
		sum += matchsticks[i]
	}
	if len(matchsticks) < 4 {
		return false
	}
	if sum%4 != 0 {
		return false
	}
	sideLength := sum / 4
	bucket := make([]int, len(matchsticks))
	return DFS(matchsticks, bucket, len(matchsticks)-1, sideLength)
}

func DFS(nums, bucket []int, index, target int) bool {
	if index < 0 {
		return true
	} else {
		for i := 0; i < 4; i++ {
			if bucket[i] == target {
				continue
			}
			if bucket[i]+nums[index] > target { // 不能放置到当前桶中
				continue
			}
			// 放置到当前桶中
			bucket[i] += nums[index]
			if DFS(nums, bucket, index-1, target) {
				// 下一个火柴可以防止到桶里
				return true
			}
			bucket[i] -= nums[index] // 将火柴从桶里拿出来
		}
		// 当前火柴任何一个桶都放不下
		return false
	}
}
