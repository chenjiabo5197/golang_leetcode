package main

import "fmt"

var result = make([]int, 0)
var array = [8][2]int{}

func main() {
	eightQueen()
}

func eightQueen() {
	for i:=0;i<8;i++ {
		getResult(i,0)
		array[0][0] = i
		array[0][1] = 0
		result = append(result, i)
	}
}

func getResult(row int, col int) {
	if len(result) == 8 {
		fmt.Println(result)
		result = []int{}
	}

	for j:=col;j<8;j++ {
		array[row][j] = 1
		result = append(result, j)
		getResult(row+1,col)
	}

}

func isSet(row , col int) {

}
