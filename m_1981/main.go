package main

import (
	"fmt"
	"math"
)

/*

*/
func main(){
	a1 := []int{1,2,9,8,7}
	//a2 := []int{4,5,6}
	//a3 := []int{7,8,9}
	var arr [][]int
	arr = append(arr, a1)
	fmt.Println(minimizeTheDifference(arr, 6))
}

func minimizeTheDifference(mat [][]int, target int) int {
	sumMap := make(map[int]bool)
	sumMap[0] = true
	for _,arr := range mat {
		tmpMap := make(map[int]bool)
		minMax := 70*70
		for _, arrn := range arr{
			for sumNum, bl := range  sumMap {
				if !bl {
					continue
				}
				total := sumNum + arrn
				if total > target && total > minMax {
					continue
				}
				if total > target && total < minMax {
					sumMap[minMax] = false
					minMax = total
				}
				tmpMap[total] = true
			}
		}
		sumMap = tmpMap
	}
	min := math.MaxFloat64
	for sumVal, bl := range sumMap {
		if !bl {
			continue
		}
		if math.Abs(float64(sumVal - target)) < min {
			min = math.Abs(float64(sumVal - target))
		}
	}
	return int(min)
}
