package main

import "fmt"

func main() {
	//nums := []int{2, 2, 3, 4}
	fmt.Println(solveNQueens(4))
}

//51. N 皇后
func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	//初始化棋盘
	// '.'为46   'Q'为81
	chessBoard := make([][]byte, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			chessBoard[i][j] = '.'
		}
	}
	for i := 0; i < len(chessBoard); i++ {
		temp, _ := getQueen(chessBoard, 0, i)
		result = append(result, temp...)
	}
	//fmt.Printf("temp=%v\n", temp)
	return result
}

func getQueen(chessBoard [][]byte, nums int, x int) ([][]string, int) {
	// 遍历
	result := make([][]string, 0)
	if x >= len(chessBoard) {
		return result, nums
	}
	for j := 0; j < len(chessBoard[0]); j++ {
		if chessBoard[x][j] == '.' && isLocate(chessBoard, x, j) {
			chessBoard[x][j] = 'Q'
			nums += 1
			tempResult, tempNums := getQueen(chessBoard, nums, x+1)
			//fmt.Printf("tempResult=%v\n", tempResult)
			result = append(result, tempResult...)
			//fmt.Printf("result=%v\n", result)
			//最后一行的皇后找到放置位置,且皇后放置完成
			if tempNums == len(chessBoard) {
				temp := make([]string, len(chessBoard))
				for k := 0; k < len(chessBoard); k++ {
					temp[k] = string(chessBoard[k])
				}
				result = append(result, temp)
				//fmt.Printf("result=%v\n", result)
			}
			nums -= 1
			chessBoard[x][j] = '.'
		}
	}

	return result, nums
}

func isLocate(chessBoard [][]byte, x, y int) bool {
	for i := 0; i < len(chessBoard); i++ { // 判断该位置能否放置皇后
		if chessBoard[i][y] == 'Q' {
			return false
		}
		if chessBoard[x][i] == 'Q' {
			return false
		}
		for j := 0; j < len(chessBoard[0]); j++ {
			if i == x && j == y {
				continue
			}
			if chessBoard[i][j] == 'Q' && (y-j == x-i || y-j == -(x-i)) {
				return false
			}
		}
	}
	return true
}
