package main

import (
	"fmt"
	"strings"
)

/*
最长公共前缀
 */

func main() {
	strs := []string{"flow", "flower", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix (strs []string) string {
	result := ""
	temp := ""
	flag := false
	if len(strs) == 1 {
		return strs[0]
	}
	for _, char := range strs[0] {
		temp += string(char)
		for i:=0;i<len(strs);i++ {
			if ok := strings.HasPrefix(strs[i], temp); !ok {
				flag = true
				break
			}
			if i == len(strs) - 1 {
				result = temp
			}
		}
		if flag {
			break
		}
	}
	return  result
}
