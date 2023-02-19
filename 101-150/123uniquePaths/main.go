package main

import "fmt"

func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n))
}

//62. 不同路径
/*
	二维数组arr，arr[i][j]表示从0,0到i，j共有arr[i][j]种走法
	arr[i][j] = arr[i-1][j] + arr[i][j-1]
*/
func uniquePaths(m int, n int) int {
	arr := make([][]int, m)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if i == 0 && j == 0 {
				arr[i][j] = 1
			} else if i == 0 && j != 0 {
				arr[i][j] = arr[i][j-1]
			} else if j == 0 && i != 0 {
				arr[i][j] = arr[i-1][j]
			} else {
				arr[i][j] = arr[i-1][j] + arr[i][j-1]
			}
		}
	}
	return arr[len(arr)-1][len(arr[0])-1]
}
