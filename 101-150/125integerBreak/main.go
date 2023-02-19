package main

import "fmt"

func main() {
	nums := 10
	fmt.Println(integerBreak(nums))
}

//343. 整数拆分
/*
	一维数组arr，arr[i] 表示将第i位数拆分之后最大的乘积
	arr[i] = max(j*(i-j), j*arr[i-j])
*/
func integerBreak(n int) int {
	arr := make([]int, n+1) //从0开始的
	arr[0] = 0
	arr[1] = 0
	for i := 2; i < len(arr); i++ {
		maxValue := 0
		for j := 1; j < i; j++ {
			maxValue = max(max(j*(i-j), j*arr[i-j]), maxValue)
		}
		arr[i] = maxValue
	}
	return arr[len(arr)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
