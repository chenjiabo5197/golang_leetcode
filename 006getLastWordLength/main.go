package main

import (
	"fmt"
	"strings"
)

/*
获取字符串中最后一个字符的长度
 */

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	result := 0
	temp := strings.Trim(s, " ")
	for i:=len(temp)-1;i>=0;i-- {
		if string(temp[i]) != " " {
			result++
		}else {
			break
		}
	}
	return result
}

