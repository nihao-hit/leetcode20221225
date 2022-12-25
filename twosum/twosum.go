package main

import (
	"fmt"
)

func main() {
	numsList := [][]int{
		{
			2,7,11,15,
		},
		{
			3,2,4,
		},
		{
			3,3,
		},
	}
	targetList := []int{
		9,6,6,
	}
	for idx, nums := range numsList {
		fmt.Println(twoSum(nums, targetList[idx]))
	}
}

func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)
	for idx, num := range nums {
		if first, ok := idxMap[num]; ok {
			return []int{first, idx}
		}
		idxMap[target-num] = idx
	}
	return []int{}
}