package main

//剑指 Offer 11. 旋转数组的最小数字
func main() {

}

// 二分解法
/*
	middle元素有3种情况，
	1、number[middle] < number[high]  则最小值在middle左侧
	2、number[middle] > number[high]  则最小值在middle右侧
	3、number[middle] == number[high] 有重复元素，故最小值左右都有可能,用线性查找
*/
func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		middle := low + (high-low)/2
		if numbers[middle] > numbers[high] {
			low = middle + 1
		} else if numbers[middle] < numbers[high] {
			high = middle
		} else {
			high--
		}
	}
	return numbers[low]
}

// 暴力解法
//func minArray(numbers []int) int {
//	temp := numbers[0]
//	for i := 1; i < len(numbers); i++ {
//		if temp > numbers[i] {
//			temp = numbers[i]
//			break
//		}
//	}
//	return temp
//}
