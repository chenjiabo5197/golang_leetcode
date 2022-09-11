package main

import (
	"errors"
	"fmt"
)

/*
图片平滑器
 */

func main() {
	arr := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println(imageSmoother(arr))
}

func imageSmoother(img [][]int) [][]int {
	row := len(img)
	col := len(img[0])
	result := make([][]int, row)
	for i:=0;i<row;i++ {
		result[i] = make([]int, col)
	}
	for i:= 0;i<row;i++ {
		for j:=0;j<col;j++ {
			num := 1
			temp := img[i][j]
			if num1, err := getNum(img, i-1, j-1); err != nil {
				temp += num1
				num++
			}
			if num1, err := getNum(img, i-1, j); err != nil {
				temp += num1
				num++
			}
			if num1, err := getNum(img, i-1, j+1); err != nil {
				temp += num1
				num++
			}
			if num1, err := getNum(img, i, j-1); err != nil {
				temp += num1
				num++
			}
			if num1, err := getNum(img, i, j+1); err != nil {
				temp += num1
				num++
			}
			if num1, err := getNum(img, i+1, j-1); err != nil {
				temp += num1
				num++
			}
			if num1, err := getNum(img, i+1, j); err != nil {
				temp += num1
				num++
			}
			if num1, err := getNum(img, i+1, j+1); err != nil {
				temp += num1
				num++
			}
			result[i][j] = temp/num
		}
	}
	return result
}

func getNum(arr [][]int, row, col int) (int, error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {

		}
	}()
	err := errors.New("test")
	return arr[row][col], err
}
