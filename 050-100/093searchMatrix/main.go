package main

import "fmt"

func main()  {
	//matrix := [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	matrix := [][]int{{11}, {2}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
}

// 74. 搜索二维矩阵 将其当做一个一维有序数组来比较
func searchMatrix(matrix [][]int, target int) bool {
	isOneDim := false
	if len(matrix) == 1 {
		isOneDim = true
	}
	cols := len(matrix[0])
	low := 0
	high := len(matrix) * len(matrix[0]) - 1
	for low <= high {
		middle := (low + high) / 2
		row := middle/cols
		if isOneDim {
			row = 0
		}
		if matrix[row][middle%cols] == target {
			return true
		}else if matrix[row][middle%cols] > target {
			high = middle - 1
		}else {
			low = middle + 1
		}
	}
	return false
}
