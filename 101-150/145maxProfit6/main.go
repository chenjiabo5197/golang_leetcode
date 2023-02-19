package main

import "fmt"

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))
}

// 714. 买卖股票的最佳时机含手续费
/*
	二维数组arr, arr[i][j]表示第i天进行j操作之后的最大利润
	j 有两种状态 0	未持有股票
				1	持有股票
*/
func maxProfit(prices []int, fee int) int {
	arr := make([][]int, len(prices))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, 2)
	}

	arr[0][0] = 0
	arr[0][1] = -prices[0]

	for i := 1; i < len(arr); i++ {
		arr[i][0] = max(arr[i-1][0], arr[i-1][1]+prices[i]-fee) //卖
		arr[i][1] = max(arr[i-1][1], arr[i-1][0]-prices[i])     //买
	}
	//fmt.Println(arr)
	return arr[len(arr)-1][0] // 因为最后一天如果持有股票会减去当天的股价，所以当前的arr[i][1]肯定小于arr[i][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
