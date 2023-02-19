package main

import "fmt"

func main() {
	num := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 5
	fmt.Println(findNumberIn2DArray(num, target))
}

//二维数组中查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		result := binarySearch(matrix[i], 0, len(matrix[i])-1, target)
		if result {
			return result
		}
	}
	return false
}

// 二分查找
func binarySearch(num []int, start, end, target int) bool {
	if start < 0 || end > len(num)-1 || start > end {
		return false
	}
	middle := (start + end) / 2
	if num[middle] == target {
		return true
	} else if num[middle] > target {
		return binarySearch(num, start, middle-1, target)
	} else {
		return binarySearch(num, middle+1, end, target)
	}
}
