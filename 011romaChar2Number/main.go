package main

import "fmt"

/*
罗马数字转整数
 */

func main() {
	s := "IX"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	table := make(map[string]int)
	table["I"] = 1
	table["V"] = 5
	table["X"] = 10
	table["L"] = 50
	table["C"] = 100
	table["D"] = 500
	table["M"] = 1000
	result := 0
	for index, value := range []rune(s) {
		if index != 0 && table[string(s[index-1])] < table[string(s[index])] {
			result += table[string(value)]
			result -= table[string(s[index-1])]
			result -= table[string(s[index-1])]
			continue
		}
		result += table[string(value)]
	}
	return result

}


