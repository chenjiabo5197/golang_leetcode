package main

import (
	"fmt"
)

/*
数据流中的中位数
 */

func main() {
	mf := Constructor()
	mf.AddNum(-1)
	fmt.Println(mf.FindMedian())
	mf.AddNum(-2)
	fmt.Println(mf.FindMedian())
	mf.AddNum(-3)
	fmt.Println(mf.FindMedian())
	mf.AddNum(-4)
	fmt.Println(mf.FindMedian())
	mf.AddNum(-5)
	fmt.Println(mf.FindMedian())
}

type MedianFinder struct {
	data []int
	length int
	isEven bool
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	cf := MedianFinder{
		data:  make([]int, 0),
		length : 0,
		isEven: true,
	}
	return cf
}


func (this *MedianFinder) AddNum(num int)  {
	//this.data = append(this.data, num)
	//sort.Ints(this.data)
	//找到新加入数字要插入的位置
	index := 0
	for index=0;index<this.length;index++ {
		if this.data[index] > num {
			break
		}
	}
	//newData := append(this.data[:index], num)
	//newData = append(newData, this.data[index:]...)
	newData := append(this.data[:index], append([]int{num}, this.data[index:]...)...)
	this.length++
	this.data = newData
	if this.isEven {
		this.isEven = false
	}else {
		this.isEven = true
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.isEven{
		return float64(this.data[(this.length/2)-1]+this.data[(this.length/2)]) / 2
	}else {
		return float64(this.data[this.length/2])
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
