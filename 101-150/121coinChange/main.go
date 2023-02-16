package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{1, 2, 5}
	target := 11
	fmt.Println(coinChange(nums, target))
}

//322. 零钱兑换
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	if amount < coins[0] {
		return -1
	}
	// arr中数字从0~amount，表示从1到amount个硬币最少的组成方法
	arr := make([]int, amount+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = -1
	}
	// 从硬币最大值的下一个位置开始遍历
	for i := coins[0]; i < len(arr); i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] == i {
				arr[i] = 1
				break
			}
		}
		minNums := math.MaxInt64
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				if arr[i-coins[j]] < minNums && arr[i-coins[j]] != -1 {
					minNums = arr[i-coins[j]]
				}
			}
		}
		if minNums != math.MaxInt64 && arr[i] == -1 {
			arr[i] = minNums + 1
		}
	}
	return arr[len(arr)-1]
}
