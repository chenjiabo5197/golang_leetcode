package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

//78. 子集
// 建立一个长度为len(nums)的0/1数组，0表示对应位置的数不在子集中，1表示在子集中,从0到2^n-1即可取出nums的所有子集
func subsets(nums []int) [][]int {
	max := int(math.Pow(2, float64(len(nums))) - 1)
	result := make([][]int, 0)
	for i := 0; i <= max; i++ {
		temp := make([]int, 0)
		for index, value := range nums {
			tempValue := i >> index
			if (tempValue & 1) > 0 { // 和1与保证当前0索引位置的数为1
				//fmt.Println(tempValue)
				temp = append(temp, value)
			}
		}
		//fmt.Println(temp)
		result = append(result, temp)
	}
	return result
}
