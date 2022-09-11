package main

import (
	"fmt"
)

/*
产生杨辉三角的切片
 */

func main(){
	n := 3
	fmt.Println(generate(n))
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i:=0;i<numRows;i++ {
		result[i] = make([]int,i+1)
		for j:=0;j<i+1;j++ {
			if i==0 {
				result[i][j] = 1
				continue
			}
			var temp1, temp2 int
			temp1 = getNum(result, i-1, j-1)
			temp2 = getNum(result, i-1, j)
			result[i][j] = temp1 + temp2
		}
	}
	return result
}

func getNum(arr [][]int, row, col int) int{
	defer func() {
		if panicErr := recover(); panicErr != nil {
		}
	}()
	temp := arr[row][col]
	return temp
}
