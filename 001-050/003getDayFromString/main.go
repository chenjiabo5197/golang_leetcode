package main

import (
	"fmt"
	"time"
)
/*
计算字符串时间是当年的第几天
 */

func main() {
	date := "2019-01-09"
	fmt.Println(dayOfYear(date))
}

func dayOfYear(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	year, _, _ := t.Date()
	newYear := time.Date(year, 1, 1, 0, 0, 0,0, time.Local)
	timeStamp1 := t.Unix()
	timeStamp2 := newYear.Unix()
	pastTime := timeStamp1-timeStamp2
	result := pastTime / (3600*24)
	return int(result)+1
}

