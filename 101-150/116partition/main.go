package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(partition(s))
}

var (
	res  [][]string
	path []string
)

//131. 分割回文串
func partition(s string) [][]string {
	res = make([][]string, 0)
	path = make([]string, 0)
	solve(0, s)
	return res
}

func solve(index int, s string) {
	if index == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		res = append(res, temp)
	} else {
		temp := ""
		for i := index; i < len(s); i++ {
			temp += string(s[i])
			if reverseString(temp) {
				path = append(path, temp)
				solve(i+1, s)
				path = path[:len(path)-1]
			}
		}
	}
}

// 是否回文串
func reverseString(data string) bool {
	temp := ""
	for i := len(data) - 1; i >= 0; i-- {
		temp += string(data[i])
	}
	return data == temp
}
