package main

import "fmt"

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(nums))
}

//122. 买卖股票的最佳时机 II
/*
	二维数组arr，arr[i][j]表示第i天进行了j的操作之后，获取的利润最大，买入时需要减去当天的价格，获取的利润为前面获取的最大利润加上当前的价格
	j的状态有2种，0  未持有不飘
				1   持有股票
	arr[i][0] = max(arr[i-1][0], arr[i-1][1]+price[i])
	arr[i][1] = max(arr[i-1][1], arr[i-1][0]-price[i])
*/

func maxProfit(prices []int) int {
	arr := make([][]int, len(prices))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, 3)
	}
	arr[0][0] = 0
	arr[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		arr[i][0] = max(arr[i-1][0], arr[i-1][1]+prices[i])
		arr[i][1] = max(arr[i-1][1], arr[i-1][0]-prices[i])
	}
	return arr[len(arr)-1][0] // 因为最后一天如果持有股票会减去当天的股价，所以当前的arr[i][1]肯定小于arr[i][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func maxProfit(prices []int) int {
//	index := 0 // 买股票的最佳时间
//	maxValue := 0
//	for i:=1;i<len(prices);i++ {
//		if prices[i] < prices[i-1] && prices[i-1] > prices[index] {
//			maxValue += prices[i-1] - prices[index]
//			index = i
//		}
//		if prices[index] > prices[i] {
//			index = i
//		}
//		if i==len(prices)-1 {  //递增数列
//			maxValue += prices[i] - prices[index]
//		}
//	}
//	return maxValue
//}
