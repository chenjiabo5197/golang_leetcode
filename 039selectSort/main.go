package main

import "fmt"

/*
选择排序
 */

func main() {
	arr := []int{1, 4, 7, 8, 2, 3, 9, 0, 5, 6}
	selectSort(arr)
	fmt.Println(arr)
}

func selectSort(data []int)  {
	for i:=0;i<len(data)-1;i++ {
		maxIndex := 0
		for j:=0;j<len(data)-i;j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			}
		}
		data[len(data)-i-1], data[maxIndex] = data[maxIndex], data[len(data)-i-1]

	}
}