package main

import (
	"fmt"
	"sort"
)

/*
最小弓箭数引爆气球
 */

func main() {
	data := [][]int{{9,12},{1,10},{4,11},{8,12},{3,9},{6,9},{6,7}}
	fmt.Println(findMinArrowShots(data))
}

type SortSlice [][]int

func (s SortSlice) Len() int {
	return len(s)
}

func (s SortSlice) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func (s SortSlice) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func findMinArrowShots(points [][]int) int {
	sort.Sort(SortSlice(points))
	fmt.Println(points)
	numsArrow := 0
	numsBall := len(points)
	flag := make([]bool, len(points))
	//初始化标志位
	for i:=0;i<len(flag);i++ {
		flag[i] = true
	}
	index := 0
	for numsBall > 0 {
		for index<len(points) && flag[index]{
			flag[index] = false
			numsBall--
			for j:=index+1;j<len(points);j++ {  //判断j位置的气球没有被打爆
				if !flag[j] {
					continue
				}
				if points[index][1] >= points[j][0] && points[index][1] <= points[j][1] {
					flag[j] = false  //该位置气球被打爆
					numsBall--
				}
			}
			numsArrow++
		}
		index++
	}
	return numsArrow
}