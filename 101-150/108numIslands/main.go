package main

import (
	"fmt"
)

func main() {
	nums := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(nums))
}

//200. 岛屿数量
func numIslands(grid [][]byte) int {
	// flag[i][j] 为false表示未被搜索过，为true表示已经被搜索过了
	flag := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		flag[i] = make([]bool, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			flag[i][j] = false
		}
	}
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !flag[i][j] {
				search(grid, i, j, flag)
				result += 1
			}
		}
	}
	return result
}

func search(lands [][]byte, x, y int, flag [][]bool) {
	// 越界
	if x >= len(lands) || x < 0 || y >= len(lands[0]) || y < 0 {
		return
	}
	if lands[x][y] == '1' && !flag[x][y] {
		flag[x][y] = true
		// 搜索顺序 上 -> 左 -> 下 -> 右
		search(lands, x-1, y, flag)
		search(lands, x, y-1, flag)
		search(lands, x+1, y, flag)
		search(lands, x, y+1, flag)
	}
}
