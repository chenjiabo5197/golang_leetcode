package main

import "fmt"

/*
跳跃游戏
 */

func main() {
	data := []int{3, 2, 1,0, 4}
	fmt.Println(canJump(data))
}

// 动态规划
//func canJump(nums []int) bool {
//	isJump := make([]bool, len(nums))
//	isJump[0] = true
//	for i:=0;i<len(isJump);i++ {
//		for j:=i+1;j<len(isJump) && j<=i+nums[i];j++ {
//			if isJump[i] {
//				isJump[j] = true
//			}
//		}
//	}
//	return isJump[len(isJump) -1]
//}

// 只计算当前位置最远能达到的下标
func canJump(nums []int) bool {
	indexJump := make([]int, len(nums))
	for i:=0;i<len(indexJump);i++ {
		indexJump[i] = nums[i] + i   //indexJump数组存放当前位置可以到达的最远的位置
	}
	index := 0   // 保存当前元素的下标
	maxJump := indexJump[0]   //保存最远距离
	for index < len(nums) && index <= maxJump {  //index <= maxJump是为了排除index在某一位置往后一点都挪动不了的情况，例，index处值为0
		if indexJump[index] > maxJump {  // 当前位置能达到的最远位置大于目前保存的最远位置
			maxJump = indexJump[index]
		}
		index++
	}
	if index == len(nums) {  //此处设置和上面index<=maxJump一样，如果index能到数组最后，那么会继续往后走，造成index==len(nums),
	// 否则index到不了数组最后一位的位置
		return true
	}
	return false
}