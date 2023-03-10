package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'A'}}
	word := "AA"
	fmt.Println(exist(board, word))
}

//剑指 Offer 12. 矩阵中的路径
func exist(board [][]byte, word string) bool {
	isUse := make([][]bool, len(board))
	for i := 0; i < len(isUse); i++ {
		isUse[i] = make([]bool, len(board[0]))
	}
	result := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if len(word) == 1 {
					return true
				}
				result = isExist(i, j, 0, board, isUse, word)
				if result {
					return result
				}
			}
		}
	}
	return result
}

func isExist(i, j, index int, board [][]byte, isUse [][]bool, word string) bool {
	if index >= len(word) {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || isUse[i][j] {
		return false
	}
	if word[index] == board[i][j] {
		isUse[i][j] = true
		if isExist(i-1, j, index+1, board, isUse, word) {
			return true
		}
		if isExist(i, j+1, index+1, board, isUse, word) {
			return true
		}
		if isExist(i+1, j, index+1, board, isUse, word) {
			return true
		}
		if isExist(i, j-1, index+1, board, isUse, word) {
			return true
		}
		isUse[i][j] = false
	}
	return false
}
