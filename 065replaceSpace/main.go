package main

import "fmt"

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}

//替换空格
func replaceSpace(s string) string {
	result := ""
	for i := range s {
		if s[i] == 32 {
			result += "%20"
		}else {
			result += string(s[i])
		}
	}
	return result
}