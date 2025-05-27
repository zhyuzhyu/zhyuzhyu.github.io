package main

import "fmt"

type RunCase struct {
	nums []int
}

var cases []RunCase

func removeDuplicates(nums []int) int {
	left := 0
	for i := range nums {
		if i+1 == len(nums) || nums[i] != nums[i+1] {
			nums[left] = nums[i]
			left++
		}
	}

	return left
}

func initCases() {
	cases = append(cases, RunCase{
		[]int{1, 1, 2},
	})

	cases = append(cases, RunCase{
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	})
}

func main() {
	initCases()

	bk := 0
	for i, c := range cases {
		if i < bk {
			continue
		}

		idx := removeDuplicates(c.nums)
		fmt.Println(c.nums, idx)
	}
}
