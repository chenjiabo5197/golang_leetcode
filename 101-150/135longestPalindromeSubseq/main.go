package main

import "fmt"

func main() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
}

//516. 最长回文子序列
func longestPalindromeSubseq(s string) int {
	text1 := s
	text2 := reverseString(s)
	arr := make([][]int, len(text1))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, len(text2))
	}
	if text1[0] == text2[0] {
		arr[0][0] = 1
	} else {
		arr[0][0] = 0
	}
	for i := 1; i < len(arr); i++ {
		if text1[i] == text2[0] {
			arr[i][0] = 1
		} else {
			arr[i][0] = arr[i-1][0]
		}
	}
	for j := 1; j < len(arr[0]); j++ {
		if text1[0] == text2[j] {
			arr[0][j] = 1
		} else {
			arr[0][j] = arr[0][j-1]
		}
	}
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[0]); j++ {
			if text1[i] == text2[j] {
				arr[i][j] = max(arr[i-1][j-1]+1, max(arr[i-1][j], arr[i][j-1]))
			} else {
				arr[i][j] = max(arr[i-1][j-1], max(arr[i-1][j], arr[i][j-1]))
			}
		}
	}
	return arr[len(arr)-1][len(arr[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverseString(s string) string {
	temp := ""
	for i := len(s) - 1; i >= 0; i-- {
		temp += string(s[i])
	}
	return temp
}
