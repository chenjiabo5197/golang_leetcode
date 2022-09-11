package main

import (
	"fmt"
	"sort"
)

/*
最小的k个数
 */

func main() {
	arr := []int{0,1,2,1}
	k := 1
	fmt.Println(getLeastNumbers(arr, k))
}

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	result := make([]int, 0)
	for i:=0;i<k;i++ {
		result = append(result, arr[i])
	}
	return result
}