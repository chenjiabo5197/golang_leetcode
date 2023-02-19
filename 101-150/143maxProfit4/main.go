package main

import "fmt"

func main() {
	k := 2
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit(k, prices))
}

// 188. 买卖股票的最佳时机 IV
/*
	三维数组arr, arr[i][j][k]表示第i天进行第j次交易k操作之后的最大利润
	k 有两种状态 0	未持有股票
				1	持有股票
	j 最多为k，表示完成了k次交易,最大为利润arr[i][k][0]
*/

func maxProfit(k int, prices []int) int {
	arr := make([][][]int, len(prices))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([][]int, k+1)
		for j := 0; j < len(arr[0]); j++ {
			arr[i][j] = make([]int, 2)
		}
	}

	for i := 0; i <= k; i++ { // 初始化,第1天的情况
		arr[0][i][0] = 0
		arr[0][i][1] = -prices[0]
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j <= k; j++ {
			if j == 0 {
				arr[i][0][0] = arr[i-1][0][0]
			} else {
				arr[i][j][0] = max(arr[i-1][j][0], arr[i-1][j-1][1]+prices[i])
			}
			arr[i][j][1] = max(arr[i-1][j][1], arr[i-1][j][0]-prices[i])
		}
	}
	//fmt.Println(arr)
	return arr[len(arr)-1][k][0] // 因为最后一天如果持有股票会减去当天的股价，所以当前的arr[i][1]肯定小于arr[i][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
