package main

import "fmt"

/*
快速排序
 */

func main() {
	arr := []int{1, 4, 7, 8, 2, 3, 9, 0, 5, 6}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(data []int, low, high int)  {
	if low >= high {
		return
	}
	indexNum := data[high]
	start := low
	end := high
	for start < end {
		for start< end && data[start] < indexNum {
			start++
		}
		data[end] = data[start]
		for start< end && data[end] > indexNum {
			end--
		}
		data[start] = data[end]
	}
	data[end] = indexNum
	fmt.Println(data)
	fmt.Println(low, end)
	quickSort(data, low, end-1)
	quickSort(data, end+1, high)
}
