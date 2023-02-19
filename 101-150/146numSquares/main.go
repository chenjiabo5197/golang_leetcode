package main

import (
	"fmt"
	"math"
)

func main() {
	n := 12
	fmt.Println(numSquares(n))
}

// 279. 完全平方数
/*
	一维数组arr, arr[i]表示组成i的最少完全平方数数量
*/
func numSquares(n int) int {
	arr := make([]int, n+1)
	arr[0] = 0
	for i := 1; i < len(arr); i++ {
		arr[i] = math.MaxInt64
	}
	for i := 0; i < len(arr); i++ {
		increase := 1
		for i-increase*increase >= 0 {
			arr[i] = min(arr[i-(increase*increase)]+1, arr[i])
			increase++
		}

	}
	return arr[len(arr)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
