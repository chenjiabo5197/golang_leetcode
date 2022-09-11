package main

import "fmt"

/*
冒泡排序
 */

func main() {
	arr := []int{1, 4, 7, 8, 2, 3, 9, 0, 5, 6}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(data []int) {
	for i:=0;i<len(data)-1;i++ {
		for j:=0;j<len(data)-i-1;j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}