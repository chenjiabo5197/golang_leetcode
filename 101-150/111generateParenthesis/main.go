package main

import "fmt"

func main() {
	//nums := []int{2, 2, 3, 4}
	fmt.Println(generateParenthesis(3))
}

var (
	res []string
)

//22. 括号生成
func generateParenthesis(n int) []string {
	res = make([]string, 0)
	DFS("", n, n)
	return res
}

func DFS(temp string, left, right int) {
	if left == 0 && right == 0 {
		res = append(res, temp)
	} else {
		if left == right {
			DFS(temp+"(", left-1, right)
		} else if left < right {
			if left > 0 {
				DFS(temp+"(", left-1, right)
			}
			DFS(temp+")", left, right-1)
		}
	}
}
