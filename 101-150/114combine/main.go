package main

import "fmt"

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}

var (
	res [][]int
	sub []int
)

//77. 组合
func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	subSlice(k, 0, n)
	return res
}

func subSlice(length, index, end int) {
	if len(sub) == length {
		temp := make([]int, len(sub))
		copy(temp, sub)
		res = append(res, temp)
	} else {
		for i := index + 1; i <= end; i++ {
			if len(sub) < length {
				sub = append(sub, i)
				subSlice(length, i, end)
			}
			sub = sub[:len(sub)-1]
		}
	}
}
