package main

import "fmt"

func main() {
	m := 16
	n := 8
	k := 4
	fmt.Println(movingCount(m, n, k)) // 15
}

//面试题13. 机器人的运动范围
func movingCount(m int, n int, k int) int {
	res := 0
	isReach := make([][]bool, m)
	for i := 0; i < len(isReach); i++ {
		isReach[i] = make([]bool, n)
	}
	isReach[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if reachPlace(i, j, isReach) {
				temp := getSum(i, j)
				if temp <= k {
					isReach[i][j] = true
					res++
				}
			}
		}
	}
	//fmt.Println(isReach)
	return res
}

func reachPlace(i, j int, isReach [][]bool) bool {
	if i <= 0 && j <= 0 {
		return isReach[i][j]
	}
	if i <= 0 {
		return isReach[i][j-1]
	}
	if j <= 0 {
		return isReach[i-1][j]
	}
	return isReach[i][j-1] || isReach[i-1][j]
}

func getSum(x, y int) int {
	hundredX := x / 100
	hundredY := y / 100
	tenX := (x % 100) / 10
	tenY := (y % 100) / 10
	oneX := x % 10
	oneY := y % 10
	res := hundredX + hundredY + tenX + tenY + oneX + oneY
	//fmt.Println(x, y, res)
	return res
}
