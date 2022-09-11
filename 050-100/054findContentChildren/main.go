package main

import (
	"fmt"
	"sort"
)

/*
分发饼干
 */

func main() {
	g := []int{1,2}
	s := []int{1,2,3}
	fmt.Println(findContentChildren(g, s))
}

// 贪心算法  怎么证明正确？
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	indexg := 0
	indexs := 0
	for indexs<len(s) && indexg<len(g) {
		if s[indexs] >= g[indexg] {
			indexs++
			indexg++
		}else {
			indexs++
		}
	}
	return indexg
}