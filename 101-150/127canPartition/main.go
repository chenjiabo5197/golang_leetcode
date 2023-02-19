package main

import "fmt"

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}

//416. 分割等和子集
/*
	0-1背包问题
	二维数组arr，横坐标为数组之和的一半，纵坐标为nums数组，arr[i][j] 表示到和为j时，nums中i元素能否选取
*/
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	arr := make([][]bool, len(nums))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]bool, target+1) // 从0到target
	}
	for i := 0; i < len(arr); i++ {
		arr[i][0] = false
	}
	for j := 0; j < len(arr[0]); j++ {
		if nums[0] == j {
			arr[0][j] = true
		} else {
			arr[0][j] = false
		}
	}
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[0]); j++ {
			if j-nums[i] >= 0 {
				arr[i][j] = arr[i-1][j] || arr[i-1][j-nums[i]]
			} else {
				arr[i][j] = arr[i-1][j]
			}
		}
	}
	return arr[len(arr)-1][len(arr[0])-1]
}
