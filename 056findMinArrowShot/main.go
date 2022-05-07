package main

import (
	"fmt"
	"sort"
)

/*
最小弓箭数引爆气球
 */

func main() {
	data := [][]int{{3,9},{7,12},{3,8},{6,8},{9,10},{2,9},{0,9},{3,9},{0,6},{2,8}}
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
	return s[i][1] < s[j][1]
}

//func findMinArrowShots(points [][]int) int {
//	sort.Sort(SortSlice(points))
//	fmt.Println(points)
//	numsArrow := 0
//	numsBall := len(points)
//	flag := make([]bool, len(points))
//	//初始化标志位
//	for i:=0;i<len(flag);i++ {
//		flag[i] = true
//	}
//	index := 0
//	for numsBall > 0 {
//		for index<len(points) && flag[index]{
//			flag[index] = false
//			numsBall--
//			for j:=index+1;j<len(points);j++ {  //判断j位置的气球没有被打爆
//				if !flag[j] {
//					continue
//				}
//				if points[index][1] >= points[j][0] && points[index][1] <= points[j][1] {
//					flag[j] = false  //该位置气球被打爆
//					numsBall--
//				}
//			}
//			numsArrow++
//		}
//		index++
//	}
//	return numsArrow
//}

// 简化思路，只考虑如果下一个,用线段的末尾进行排序，若当前弓箭的小于下一段的起点，则需要新加一个弓箭
func findMinArrowShots(points [][]int) int {
	sort.Sort(SortSlice(points))
	nowArrow := points[0][1] //当前弓箭位置
	nums := 1
	for i:=1;i<len(points);i++ {
		if nowArrow < points[i][0] {
			nowArrow = points[i][1]
			nums++
		}
	}
	return nums
}