package main

import (
	"fmt"
)

/*
滑动窗口最大值
 */

func main() {

	nums := []int{1,-1}
	k := 1
	fmt.Println(maxSlidingWindow(nums, k))
}

// 使用队列进行操作，只需要遍历一次切片即可
func maxSlidingWindow(nums []int, k int) []int {
	queue := &MyQueue{
		data:      make([]int, 0),
		headPoint: -1,
		footPoint: -1,
	}
	result := make([]int, 0)
	if len(nums) == 0 || k == 0 {
		return result
	}
	//先向滑动窗口中添加元素，等到达k时再取值
	for i:=0;i<k;i++ {
		if queue.empty() {
			queue.push(nums[i])
		}else {
			//每添加一个元素均要与队尾元素进行比较，若大于队尾元素，则移除队尾元素，再与队尾元素比较，直至队列为空或者队尾元素大于当前元素，将其加入队尾
			for {
				if queue.empty() || queue.data[0] >= nums[i] {  //此处要将等于队尾的也添加进去，防止出现相同大小的元素
					queue.push(nums[i])
					break
				}
				queue.remove()
			}
		}
	}
	result = append(result, queue.data[queue.headPoint])

	//从此开始，每滑动一次都产生一个结果
	for i:=k;i<len(nums);i++ {
		//每次滑动前先判断当前队列中最大的数字是否要出窗口了
		if queue.data[queue.headPoint] == nums[i-k] {
			queue.pop()
		}
		for {
			if queue.empty() || queue.data[0] >= nums[i] {  //此处要将等于队尾的也添加进去，防止出现相同大小的元素
				queue.push(nums[i])
				break
			}
			queue.remove()
		}
		result = append(result, queue.data[queue.headPoint])

	}
	return result
}

type MyQueue struct {
	data []int
	headPoint int
	footPoint int
}

func (queue *MyQueue) push (value int) {
	tempSlice := []int{value}
	queue.data = append(tempSlice, queue.data...)
	queue.headPoint++
}

//从队列头部移除元素
func (queue *MyQueue) pop () {
	if queue.empty() {
		return
	}
	queue.data = queue.data[:len(queue.data)-1]
	queue.headPoint--
}

//从队列尾部移除元素
func (queue *MyQueue) remove()  {
	if queue.empty() {
		return
	}
	queue.data = queue.data[1:]
	queue.headPoint--
}

func (queue *MyQueue) peak () int{
	if queue.empty() {
		return 0
	}
	return queue.data[queue.headPoint]
}

func (queue *MyQueue) empty () bool {
	return len(queue.data) == 0
}

//func (queue *MyQueue) getMaxValue() int {
//	tempSlice := make([]int, len(queue.data))
//	copy(tempSlice, queue.data)
//	sort.Ints(tempSlice)
//	return tempSlice[len(tempSlice)-1]
//}
