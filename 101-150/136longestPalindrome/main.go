package main

import "fmt"

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
}

//5. 最长回文子串
/*
	二维数组arr，arr[i][j]表示s中从i到j位置的元素是否为回文子串
	i<=j
*/
func longestPalindrome(s string) string {
	arr := make([][]bool, len(s))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]bool, len(s))
	}
	arr[len(arr)-1][len(arr[0])-1] = true
	maxLength := string(s[len(s)-1])
	for i := len(arr) - 2; i >= 0; i-- {
		for j := i; j < len(arr[0]); j++ {
			if j == i {
				arr[i][j] = true // i和j相等，证明只有一个元素，所以肯定为回文子串
			} else {
				if s[i] != s[j] {
					arr[i][j] = false
				} else {
					if i+1 <= j-1 { // i小于等于j
						arr[i][j] = arr[i+1][j-1]
					} else {
						arr[i][j] = true
					}
				}
			}
			if arr[i][j] && len(maxLength) < j-i+1 {
				maxLength = s[i : j+1]
			}
		}
	}
	return maxLength
}
