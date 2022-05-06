package main

import "fmt"

/*
部分排序
 */

func main() {
	data := []int{1, 3, 9, 7, 5}
	fmt.Println(subSort(data))
}

// 方法错误，没有考虑要排序数组里面的值
//func subSort(array []int) []int {
//	result := []int{-1, -1}
//	for i:=0;i<len(array);i++ {
//		for j:=len(array)-1;j>i;j-- {
//			if array[j] == array[i] {
//				for k:=i;k<j;k++ {
//					if array[k] != array[i] {
//						result[0] = i
//						result[1] = j
//						fmt.Println(array[i])
//						fmt.Println(array[j])
//						return result
//					}
//				}
//			}else if array[j] < array[i] {
//				result[0] = i
//				result[1] = j
//				fmt.Println(array[i],)
//				fmt.Println(array[j],)
//				return result
//			}
//		}
//	}
//	return result
//}

// 两次遍历，从前，从后各一次，分别找到索引
// 注意，从前面开始遍历，找到的为n，即n的右边的数字都比n左边数字的最大值大，且是有序的
//		从后面开始遍历，找到的为m
func subSort(array []int) []int {
	result := []int{-1, -1}
	if len(array) <= 1 {
		return result
	}
	m := -1
	max := array[0]
	for i:=1;i<len(array);i++ {
		if array[i] < max {
			m = i
		}else {
			max = array[i]
		}
	}
	if m == -1 {
		return result
	}
	n := -1
	min := array[len(array)-1]
	for i:=len(array)-2;i>=0;i-- {
		if array[i] > min {
			n = i
		}else {
			min = array[i]
		}
	}
	result[0] = n
	result[1] = m
	return result
}
