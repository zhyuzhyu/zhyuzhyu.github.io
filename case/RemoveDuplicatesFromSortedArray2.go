package main

import "fmt"

type RunCase struct {
	nums []int
}

var cases []RunCase

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return l
	}

	quick, slow := 2, 2
	for ; quick < len(nums); quick++ {
		if nums[quick] != nums[slow-2] {
			nums[slow] = nums[quick]
			slow++
		}
	}

	return slow
}

func initCases() {
	cases = append(cases, RunCase{
		[]int{1, 1, 1, 2, 2, 2, 3},
	})

	cases = append(cases, RunCase{
		[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
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
