package main

import "fmt"

func main() {
	nums := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(nums))
}

// 123. 买卖股票的最佳时机 III
/*
	三维数组arr, arr[i][j][k]表示第i天进行第j次交易k操作之后的最大利润
	k 有两种状态 0	未持有股票
				1	持有股票
	j 最多为2，表示完成了两次交易,最大为利润arr[i][2][0]
*/

func maxProfit(prices []int) int {
	arr := make([][][]int, len(prices))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([][]int, 3)
		for j := 0; j < len(arr[0]); j++ {
			arr[i][j] = make([]int, 2)
		}
	}
	arr[0][0][0] = 0
	arr[0][0][1] = -prices[0]
	arr[0][1][0] = 0
	arr[0][1][1] = -prices[0]
	arr[0][2][0] = 0
	for i := 1; i < len(arr); i++ {
		arr[i][0][0] = arr[i-1][0][0]
		//第一次买入和卖出
		arr[i][0][1] = max(arr[i-1][0][1], arr[i-1][0][0]-prices[i])
		arr[i][1][0] = max(arr[i-1][1][0], arr[i-1][0][1]+prices[i])

		//第二次买入和卖出
		arr[i][1][1] = max(arr[i-1][1][1], arr[i-1][1][0]-prices[i])
		arr[i][2][0] = max(arr[i-1][2][0], arr[i-1][1][1]+prices[i])
	}
	return arr[len(arr)-1][2][0] // 因为最后一天如果持有股票会减去当天的股价，所以当前的arr[i][1]肯定小于arr[i][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
