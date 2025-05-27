package main

import "fmt"

type RunCase struct {
	nums []int
	val  int
}

var cases []RunCase

func removeElement(nums []int, val int) int {
	left := 0
	for i := range nums {
		if nums[i] != val {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

func initCases() {
	cases = append(cases, RunCase{
		[]int{3, 2, 2, 3},
		3,
	})

	cases = append(cases, RunCase{
		[]int{},
		0,
	})
}

func main() {
	initCases()

	bk := 0
	for i, c := range cases {
		if i < bk {
			continue
		}

		idx := removeElement(c.nums, c.val)
		fmt.Println(c.nums, idx)
	}
}
