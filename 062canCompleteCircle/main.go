package main

import "fmt"

/*
加油站，计算是否能绕圈
 */

func main() {
	gas := []int{1,2,3,4,5,5,70}
	cost := []int{2,3,4,3,9,6,2}
	fmt.Println(canCompleteCircuit(gas, cost))
}

/*
可以解题，耗时过长
	分析：一直坚持每次找到一个点都去尝试循环遍历所有点以证明从该点出发可以遍历完所有点，
	且未考虑如果从i点出发到达不了j点，则下次直接从j点出发，不需要去尝试i与j之间的点
*/
//func canCompleteCircuit(gas []int, cost []int) int {
//	leftGas := 0
//	for i:=0;i<len(gas);i++ {
//		leftGas += gas[i] - cost[i]
//	}
//	if leftGas < 0 {
//		return -1
//	}
//	if len(gas) == 1 {
//		if gas[0] >= cost[0] {
//			return 0
//		}else {
//			return -1
//		}
//	}
//	for i:=0;i<len(gas);i++ {
//		if gas[i] >= cost[i] {  //从此处开始
//			index := i+1  //起始点的坐标
//			if index == len(gas) {
//				index = 0
//			}
//			leftGas = gas[i] - cost[i] //剩余的汽油
//			for index != i {
//				leftGas += gas[index] - cost[index]
//				if leftGas < 0 {
//					break
//				}
//				index++
//				if index == len(gas) {
//					index = 0
//				}
//			}
//			if index == i && leftGas >= 0 {
//				return index
//			}
//		}
//	}
//	return -1
//}

/*
证明：如果gas列表之和>=cost列表之和，则一定存在一个点，从其出发可以遍历完所有点
	当n=2时，则AB两点一定可以找到一个点的gas>=cost,则从其出发可以到达另外一个点，然后根据gas之和>=cost之和可以得出，可以从到达点回到起始点
	当n>2时，可以将其看为是n-1个点的集合和1个点两个点之间的路程，然后将n-1个点一直拆解到2个点，
	综上，假设得证
 */
func canCompleteCircuit(gas []int, cost []int) int {
	leftGas := 0
	for i:=0;i<len(gas);i++ {
		leftGas += gas[i] - cost[i]
	}
	if leftGas < 0 {
		return -1
	}
	index := 0  //遍历的下标
	startIndex := 0   //保存起始点的下标使用
	leftGas = 0  //进入循环前初始化剩余量
	for index < len(gas) {
		leftGas += gas[index] - cost[index]
		if leftGas < 0 {
			leftGas = 0 //无法到达下一个点，重新初始化剩余汽油量
			index++
			startIndex = index  //从未到达的下一个点出发
		}else {
			index++
		}
	}
	return startIndex
}


