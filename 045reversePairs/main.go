package main

import (
	"fmt"
)

/*
数组中的逆序对
 */

func main() {
	data := []int{}
	fmt.Println(reversePairs(data))
}

//超时了
//func reversePairs(nums []int) int {
//	tempSlice := make([]int, len(nums))
//	copy(tempSlice, nums)
//	sort.Ints(tempSlice)
//	res := 0
//	for i:=0;i<len(nums);i++ {
//		for j:=0;j<len(tempSlice);j++ {
//			if nums[i] == tempSlice[j] {
//				res += j
//				//将下标为j的数据从tempSlice中删除
//				tempSlice = append(tempSlice[:j], tempSlice[j+1:]...)
//				break
//			}
//		}
//	}
//	return res
//}

//使用归并排序解决，在合并两个切片时统计数量
func reversePairs(nums []int) int {
	_, result := mergeSort(nums)
	return result
}

func mergeSort(data []int) ([]int, int) {
	if len(data) <= 1{
		return data, 0
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
	arr1, res1 := mergeSort(data1)
	arr2, res2 := mergeSort(data2)
	result, res := mergeSlice(arr1, arr2, res1+res2)   //用排好序的切片进行合成
	//result := mergeSlice([]int{7}, []int{8, 2})
	//fmt.Println(data1)
	//fmt.Println(data2)
	//fmt.Println(result)
	return result, res
}
func mergeSlice(data1, data2 []int, res int) ([]int, int) {
	length := len(data1)+len(data2)
	data3 := make([]int, length)
	index1 := 0
	index2 := 0
	for i:=0;i<length;i++ {
		if index1 < len(data1) && index2 < len(data2) {
			if  data1[index1] > data2[index2] {
				data3[i] = data2[index2]
				res += len(data1)-index1
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
	return data3, res
}
