package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}

//64. 最小路径和
/*
	二维数组arr,arr[i][j]表示移动到当前位置的最小路径和
	要到i，j位置，只能从i-1，j位置或者i，j-1位置移动过来，所以arr[i][j] = min(arr[i-1][j],arr[i][j-1])
*/
func minPathSum(grid [][]int) int {
	//arr数组比grid数组横纵坐标各大一，为了避免边界越界
	arr := make([][]int, len(grid)+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, len(grid[0])+1)
	}
	for i := 0; i < len(arr); i++ {
		arr[i][0] = math.MaxInt64
	}
	for j := 0; j < len(arr[0]); j++ {
		arr[0][j] = math.MaxInt64
	}
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[0]); j++ {
			if i == 1 && j == 1 {
				arr[i][j] = grid[0][0]
			} else {
				arr[i][j] = int(math.Min(float64(arr[i-1][j]), float64(arr[i][j-1]))) + grid[i-1][j-1]
			}
		}
	}
	return arr[len(arr)-1][len(arr[0])-1]
}
