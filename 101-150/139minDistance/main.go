package main

import (
	"fmt"
	"math"
)

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}

//72. 编辑距离
/*
 arr二维数组，arr[i][j]表示到word1的第i位，word2的第j位所进行的最少操作
arr[i][j] = min(arr[i-1][j], arr[i-1][j-1],arr[i][j-1]),
arr[i-1][j] 表示删除word1第i位字母
arr[i-1][j-1] 表示替换word1第i位字母
arr[i][j-1] 表示增加word1第i位字母
*/
func minDistance(word1 string, word2 string) int {
	arr := make([][]int, len(word1)+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(arr); i++ {
		arr[i][0] = i // 空字符与字符相比，删除字符
	}
	for j := 0; j < len(arr[0]); j++ {
		arr[0][j] = j // 空字符与字符相比，删除字符
	}
	arr[0][0] = 0 // 空字符与空字符不用进行任何操作
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[0]); j++ {
			minValue := int(math.Min(float64(arr[i-1][j-1]), math.Min(float64(arr[i-1][j]), float64(arr[i][j-1]))))
			if word1[i-1] == word2[j-1] { // 相等字符不需要进行额外操作,直接取上一个字符比较处的值即可
				arr[i][j] = arr[i-1][j-1]
			} else {
				arr[i][j] = minValue + 1
			}
		}
	}
	return arr[len(arr)-1][len(arr[0])-1]
}
