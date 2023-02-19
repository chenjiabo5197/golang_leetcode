package main

import (
	"fmt"
	"math"
)

func main() {
	nums := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(nums))
}

// 120. 三角形最小路径和
/*
	二维数组arr,arr[i][j] 表示到i，j位置的最小路径
	arr[i][j] = min(arr[i-1][j], arr[i-1][j-1])
*/
func minimumTotal(triangle [][]int) int {
	arr := make([][]int, len(triangle))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, len(triangle[i]))
	}
	arr[0][0] = triangle[0][0]
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if j == 0 { // 没有斜上方数字
				arr[i][j] = arr[i-1][j]
			} else if j == len(arr[i])-1 { // 没有正上方数字
				arr[i][j] = arr[i-1][j-1]
			} else {
				arr[i][j] = min(arr[i-1][j], arr[i-1][j-1])
			}
			arr[i][j] += triangle[i][j]
		}
	}
	minPath := math.MaxInt64
	for j := 0; j < len(arr[len(arr)-1]); j++ {
		minPath = min(minPath, arr[len(arr)-1][j])
	}
	return minPath
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
