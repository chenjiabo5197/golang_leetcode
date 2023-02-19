package main

import "fmt"

func main() {
	nums := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	fmt.Println(calculateMinimumHP(nums))
}

//174. 地下城游戏
/*
	三维数组arr，arr[i][j]表示到i，j位置需要的最小生命值
	该数组从右下向左上填充
*/
func calculateMinimumHP(dungeon [][]int) int {
	arr := make([][]int, len(dungeon))
	for i := 0; i < len(dungeon); i++ {
		arr[i] = make([]int, len(dungeon[0]))
	}
	arr[len(arr)-1][len(arr[0])-1] = max(1-dungeon[len(dungeon)-1][len(dungeon[0])-1], 1)

	for i := len(arr) - 2; i >= 0; i-- {
		arr[i][len(arr[0])-1] = max(1, arr[i+1][len(arr[0])-1]-dungeon[i][len(arr[0])-1])
	}
	for j := len(arr[0]) - 2; j >= 0; j-- {
		arr[len(arr)-1][j] = max(1, arr[len(arr)-1][j+1]-dungeon[len(arr)-1][j])
	}
	for i := len(arr) - 2; i >= 0; i-- {
		for j := len(arr[0]) - 2; j >= 0; j-- {
			arr[i][j] = max(min(arr[i+1][j], arr[i][j+1])-dungeon[i][j], 1)
		}
	}
	return arr[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
