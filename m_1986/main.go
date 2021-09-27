package main

import (
	"fmt"
	"sort"
)

func main()  {
	//task := []int{1,5,7,10,3,8,4,2,6,2}
	task := []int{1,2,3}
	//task := []int{2,3,3,4,4,4,6,7,8,9,10}
	fmt.Println(minSessionsV2(task, 3))
}

func minSessionsV2(tasks []int, sessionTime int) int {
	taskByte := 1 << len(tasks)
	var spendArr []int
	for i := 1; i <= taskByte; i++ {
		spendArr = append(spendArr, len(tasks) + 1)
	}
	for i := 1 ;i < taskByte ; i++ {
		bt := i
		spend := 0
		index := 0
		for bt != 0 {
			if bt & 1 == 1 {
				spend += tasks[index]
			}
			bt = bt >> 1
			index++
		}
		if spend <= sessionTime {
			spendArr[i] = 1
		}
	}
	//fmt.Println(spendArr)

	for i := 1; i < taskByte; i++ {
		if spendArr[i] == 1 {
			continue
		}
		for j := i; j > 0; j = (j-1) & i {
			fmt.Println(j, i ^ j)
			if spendArr[i] > spendArr[j] + spendArr[i ^ j] {
				spendArr[i] = spendArr[j] + spendArr[i ^ j]
			}
		}
	}
	return spendArr[taskByte - 1]
}


func minSessions(tasks []int, sessionTime int) int {
	tar := 1 << len(tasks)
	for i := 0; i < len(tasks); i++ {
		tar = SetIndex(tar, i, len(tasks))
	}
	p := &Const{
		TimeLen: sessionTime,
		min : len(tasks),
		TotalNum: tar,
	}
	fmt.Println(p)
	sort.Ints(tasks)

	Sessions(tasks, 1 << len(tasks), sessionTime, 1, p)
	return p.min
}
type Const struct {
	TimeLen int
	min int
	TotalNum int
}
func Sessions(tasks []int,byteNum int,sessionTime int, current int,p *Const)  {
	//fmt.Println(byteNum, sessionTime, current)
	if byteNum == p.TotalNum {
		if current < p.min {
			p.min = current
		}
		return
	}

	for i := len(tasks) - 1 ; i >= 0; i-- {
		if IsTrue(byteNum, i, len(tasks)) {
			continue
		}
		NewNum := SetIndex(byteNum, i, len(tasks))
		if tasks[i] <= sessionTime {
			Sessions(tasks, NewNum, sessionTime - tasks[i], current, p)
		} else {
			Sessions(tasks, NewNum, p.TimeLen - tasks[i]  , current + 1, p)
		}
	}
}

//index 数组下表 len 数组长度
func IsTrue(tar int, index int, len int) bool {
	if tar & (1 << (len - index - 1)) == 0 {
		return false
	}
	return true
}

func SetIndex(tar int, index int, len int) int {
	return tar | (1 << (len - index - 1))
}