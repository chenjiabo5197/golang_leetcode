package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(cuttingRope(120))
}

//剑指 Offer 14- II. 剪绳子 II
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
/*
	一维数组arr，arr[i]表示第i位数字时最大值
	arr[i] = max(arr[i], max(j*arr[i-j], j*(i-j)))   953271190
*/
func cuttingRope(n int) int {
	arr := make([]*big.Int, n+1) // 从0开始
	for i := 0; i < len(arr); i++ {
		arr[i] = new(big.Int)
	}
	for i := 2; i < len(arr); i++ {
		maxValue := new(big.Int)
		for j := 1; j < i; j++ {
			temp1 := new(big.Int)
			temp2 := new(big.Int)
			temp1.Mul(big.NewInt(int64(j)), arr[i-j])
			temp2.Mul(big.NewInt(int64(j)), big.NewInt(int64(i-j)))
			maxValue = max(maxValue, max(temp1, temp2))
		}
		arr[i] = maxValue
	}
	var result big.Int
	return int(result.Mod(arr[n], big.NewInt(1000000007)).Int64())
}

func max(a, b *big.Int) *big.Int {
	maxValue := a.Cmp(b)
	if maxValue == 1 {
		return a
	}
	return b
}
