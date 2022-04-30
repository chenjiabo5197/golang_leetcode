package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
洗牌算法
 */

func main() {
	arr := [6]int{0, 0, 0, 1, 0, 0}
	n := len(arr)
	result := make(map[int]int)
	result[0] = 0
	result[1] = 0
	result[2] = 0
	result[3] = 0
	result[4] = 0
	result[5] = 0

	for j:=0;j<1000;j++ {
		for i:=0;i<n-1;i++ {
			temp := randInt(i, n-1)
			swap(&arr[i], &arr[temp])
		}

		for i:=0;i<n;i++ {
			if arr[i] == 1 {
				result[i] ++
			}
		}
	}

	fmt.Println(result)

}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(max-min+1)+min
	return result
}

func swap(temp1, temp2 *int) {
	temp3 := *temp1
	*temp1 = *temp2
	*temp2 = temp3
}
