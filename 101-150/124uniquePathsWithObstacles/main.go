package main

import "fmt"

func main() {
	nums := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(nums))
}

//63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	arr := make([][]int, len(obstacleGrid))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, len(obstacleGrid[0]))
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				arr[i][j] = -1
			} else {
				if i == 0 && j == 0 {
					arr[i][j] = 1
				} else if i == 0 && j != 0 {
					arr[i][j] = arr[i][j-1]
				} else if j == 0 && i != 0 {
					arr[i][j] = arr[i-1][j]
				} else {
					if arr[i-1][j] != -1 && arr[i][j-1] != -1 {
						arr[i][j] = arr[i-1][j] + arr[i][j-1]
					} else if arr[i-1][j] != -1 {
						arr[i][j] = arr[i-1][j]
					} else if arr[i][j-1] != -1 {
						arr[i][j] = arr[i][j-1]
					} else {
						arr[i][j] = -1
					}
				}
			}
		}
	}
	if arr[len(arr)-1][len(arr[0])-1] == -1 {
		return 0
	}
	return arr[len(arr)-1][len(arr[0])-1]
}
