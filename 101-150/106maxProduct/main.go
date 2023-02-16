package main

import "fmt"

func main() {
	nums := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println(maxProduct(nums))
}

//318. 最大单词长度乘积
func maxProduct(words []string) int {
	bitNums := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		word := words[i]
		bitNum := 0
		for j := 0; j < len(word); j++ {
			temp := 1 << ('z' - word[j])
			//fmt.Println(temp)
			bitNum = bitNum | temp
		}
		bitNums[i] = bitNum
	}
	max := 0
	for i := 0; i < len(bitNums); i++ {
		for j := i + 1; j < len(bitNums); j++ {
			if bitNums[i]&bitNums[j] == 0 {
				tempMax := len(words[i]) * len(words[j])
				if tempMax > max {
					max = tempMax
				}
			}
		}
	}
	return max
}
