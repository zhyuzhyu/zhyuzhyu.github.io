package main

import (
	"fmt"
	"math"
)

func singleNumber(nums []int) int {
	if len(nums) <= 0 {
		return math.MinInt32
	}
	ret := nums[0]
	for _, v := range nums[1:] {
		ret ^= v
	}
	return ret
}

// [2,2,1], [4,1,2,1,2], [1]
func main() {
	arr := []int{2, 2, 1}
	arr = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(arr))
}
