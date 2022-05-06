package main

import (
	"fmt"
)

/*
计算右侧小于当前元素的个数
 */

func main() {
	arr := []int{5, 2, 6, 1}
	fmt.Println(countSmaller(arr))
}

// 暴力解法会超时
//func countSmaller(nums []int) []int {
//	result := make([]int, len(nums))
//	for i:=0;i<len(nums);i++ {
//		count := 0
//		for j:=i+1;j<len(nums);j++ {
//			if nums[j] < nums[i] {
//				count++
//			}
//		}
//		result[i] = count
//	}
//	return result
//}

// 使用归并排序来解，类似于寻找逆序对
func countSmaller(nums []int) []int {
	//先初始化一个数组，用于保存当前数组的索引,目标数组变化，索引数组也随之变化
	indexSlice := make([]int, len(nums))
	resultSlice := make([]int, len(nums))  //结果数组，储存结果
	for i:=0;i<len(indexSlice);i++ {
		indexSlice[i] = i
	}
	_, _, result := mergeSort(nums, indexSlice, resultSlice)
	fmt.Println(result)
	return result
}

func mergeSort(nums, indexSlice, resultSlice []int) ([]int, []int, []int) {
	if len(nums) <= 1 {
		return nums, indexSlice, resultSlice
	}
	mid := len(nums) / 2
	data1, index1, resultSlice := mergeSort(nums[:mid], indexSlice[:mid], resultSlice)
	data2, index2, resultSlice := mergeSort(nums[mid:], indexSlice[mid:], resultSlice)
	data3, index, result := mergeSlice(data1, data2, index1, index2, resultSlice)
	fmt.Println(data1, data2, data3, index, result)
	return data3, index, result
}

//合并两个数组
func mergeSlice(data1, data2, indexSlice1, indexSlice2, resultSlice []int) ([]int, []int, []int) {
	data := make([]int, len(data1)+len(data2))
	index := make([]int, len(data))
	index1 := 0
	index2 := 0
	for i:=0;i<len(data);i++ {
		if index1 < len(data1) && index2 < len(data2) {
			if data1[index1] > data2[index2] {
				data[i] = data2[index2]
				index[i] = indexSlice2[index2]
				index2++
			}else {
				data[i] = data1[index1]
				index[i] = indexSlice1[index1]
				resultSlice[index[i]] += index2
				index1++
			}
		}else if index1 >= len(data1) {   //data1的数字已经全部取完
			data[i] = data2[index2]
			index[i] = indexSlice2[index2]
			index2++
		}else {  //data1
			data[i] = data1[index1]
			index[i] = indexSlice1[index1]
			resultSlice[indexSlice1[index1]] += len(data2)
			index1++
		}
	}
	return data, index, resultSlice
}