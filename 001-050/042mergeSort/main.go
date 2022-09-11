package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 7, 8, 2, 3, 9, 0, 5, 6}
	_ = mergeSort(arr)
	//fmt.Println(data)
}

func mergeSort(data []int) []int {
	if len(data) == 1{
		return data
	}
	data1 := make([]int, 0)
	data2 := make([]int, 0)
	half := len(data) / 2
	for i:=0;i<len(data);i++ {
		if i<half {
			data1 = append(data1, data[i])
		}else {
			data2 = append(data2, data[i])
		}
	}
	arr1 := mergeSort(data1)
	arr2 := mergeSort(data2)
	result := mergeSlice(arr1, arr2)   //用排好序的切片进行合成
	//result := mergeSlice([]int{7}, []int{8, 2})
	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(result)
	return result
}
func mergeSlice(data1, data2 []int) []int {
	length := len(data1)+len(data2)
	data3 := make([]int, length)
	index1 := 0
	index2 := 0
	for i:=0;i<length;i++ {
		if index1 < len(data1) && index2 < len(data2) {
			if  data1[index1] > data2[index2] {
				data3[i] = data2[index2]
				index2++
			}else {
				data3[i] = data1[index1]
				index1++
			}
		}else if index1 >= len(data1) {
			data3[i] = data2[index2]
			index2++
		}else {
			data3[i] = data1[index1]
			index1++
		}

	}
	return data3
}
