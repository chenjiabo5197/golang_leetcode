package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(nums))
}

//121. 买卖股票的最佳时机
//一维数组arr，arr[i]表示第i天获取的最大利润
func maxProfit(prices []int) int {
	arr := make([]int, len(prices))
	buyValue := prices[0] // 买股票的最低值
	for i := 1; i < len(prices); i++ {
		buyValue = int(math.Min(float64(buyValue), float64(prices[i])))
		arr[i] = int(math.Max(float64(prices[i]-buyValue), float64(arr[i-1])))
	}
	return arr[len(arr)-1]
}

//func maxProfit(prices []int) int {
//	index := 0 // 买股票的最佳时间
//	maxValue := 0
//	for i:=1;i<len(prices);i++ {
//		if prices[i] - prices[index] > maxValue {
//			maxValue = prices[i] - prices[index]
//		}
//		if prices[index] > prices[i] {
//			index = i
//		}
//	}
//	return maxValue
//}
