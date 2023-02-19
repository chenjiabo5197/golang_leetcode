package main

import "fmt"

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change(amount, coins))
}

//518. 零钱兑换 II
/*
	二维数组arr,横坐标为要凑的钱数目，纵坐标为硬币的种类
arr[i][j] = arr[i-1][j] + arr[i][j-coin[i]]
即第i个硬币凑够j数量的钱的方案，等于第i-1个硬币凑够j数量的方案 加上 第i个硬币凑够j-coins[i]数量的方案
*/
func change(amount int, coins []int) int {
	arr := make([][]int, len(coins)+1)
	for i := 0; i < len(arr); i++ {
		// 横坐标从0开始，到amount截止
		arr[i] = make([]int, amount+1)
	}
	for i := 0; i < len(arr); i++ {
		arr[i][0] = 1 // 第i个硬币凑够0数量的钱只有1种方案
	}
	for j := 0; j < len(arr[0]); j++ {
		arr[0][j] = 0 // 第0个硬币凑够j数量的钱只有0种方案
	}
	for j := 1; j < len(arr[0]); j++ {
		for i := 1; i < len(arr); i++ {
			if j-coins[i-1] >= 0 { // 减去第coins[i-1]个硬币之后，硬币的面额还大于大于0,i-1是因为加了一行coins[0]的硬币，其面额为0
				arr[i][j] = arr[i-1][j] + arr[i][j-coins[i-1]]
			} else { // 减去第coins[i]个硬币之后，硬币的面额小于0，则不能取当前面值的硬币
				arr[i][j] = arr[i-1][j]
			}
		}
	}
	return arr[len(arr)-1][len(arr[0])-1]
}
