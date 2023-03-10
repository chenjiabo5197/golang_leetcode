package main

import (
	"encoding/json"
	"fmt"
)

/*
	输入：[{"name":"test1"},{"name":"test1"}{"name":"test2"}{"name":"test2"}{"name":"test1"}]
	输出：[{"name":"test1","count":2},{"name":"test1","count":"0"}{"name":"test2","count":2}{"name":"test2","count":0}{"name":"test1",,"count":1}]
*/

type Single struct {
	Name string `json:"name"`
}

type NewSingle struct {
	Single
	Count int `json:"count"`
}

func main() {
	sl := make([]Single, 0)
	nsl := make([]NewSingle, 0)
	s := "[{\"name\":\"test1\"},{\"name\":\"test1\"},{\"name\":\"test2\"},{\"name\":\"test2\"},{\"name\":\"test1\"}]"
	err := json.Unmarshal([]byte(s), &sl)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//fmt.Println(sl)
	//preKey := ""
	//for i := 0; i < len(sl); i++ {
	//	count := 1
	//	for j := i + 1; j < len(sl); j++ {
	//		if sl[i].Name == sl[j].Name {
	//			count++
	//		} else {
	//			break
	//		}
	//	}
	//	ns := NewSingle{}
	//	ns.Name = sl[i].Name
	//	if preKey == sl[i].Name {
	//		count = 0
	//	}
	//	ns.Count = count
	//	nsl = append(nsl, ns)
	//	preKey = sl[i].Name
	//}

	isReset := false
	for i := 0; i < len(sl); i++ {
		number := 1
		ns := NewSingle{}
		ns.Name = sl[i].Name
		if i != 0 && sl[i].Name == sl[i-1].Name {
			isReset = true
			number++
		} else {
			isReset = false
		}
		if isReset {
			number = 0
		}
		ns.Count = number
		nsl = append(nsl, ns)
	}

	fmt.Println(nsl)
}
