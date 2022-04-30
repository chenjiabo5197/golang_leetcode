package main

import (
	"fmt"
	"math"
)

/*
求最小下降路径和 相邻的行不能在同一列
 */

func main() {
	arr := [][]int{{2,2,1,2,2},{2,2,1,2,2},{2,2,1,2,2},{2,2,1,2,2},{2,2,1,2,2}}
	fmt.Println(minFallingPathSum(arr))
}

func minFallingPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	result := make([][]int, row)
	for i:=0;i<row;i++{
		result[i] = make([]int, col)
	}
	for i:=0;i<row;i++ {
		for j:=0;j<col;j++ {
			if i == 0{
				result[i][j] = grid[i][j]
				continue
			}
			temp := math.MaxInt32
			for k:=0;k<col;k++ {
				if j != k {
					temp2 := result[i-1][k] + grid[i][j]
					if temp > temp2 {
						temp = temp2
					}
				}
			}
			result[i][j] = temp
		}
	}
	temp := result[row-1][0]
	for i:=0;i<col;i++ {
		if temp>result[row-1][i] {
			temp = result[row-1][i]
		}
	}
	return temp
}

