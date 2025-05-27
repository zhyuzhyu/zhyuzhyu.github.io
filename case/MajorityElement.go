package main

import "fmt"

type RunCase struct {
	nums []int
}

var cases []RunCase

// 投票算法
func majorityElement(nums []int) int {
	var candi, cnt int

	for _, v := range nums {
		if cnt == 0 {
			candi = v
			cnt = 1
		} else if v == candi {
			cnt++
		} else {
			cnt--
		}
	}

	return candi
}

func initCases() {
	cases = append(cases, RunCase{
		[]int{2, 2, 1, 1, 1, 2, 2},
	})

	cases = append(cases, RunCase{
		[]int{3, 3, 4},
	})
}

func main() {
	initCases()

	bk := 0
	for i, c := range cases {
		if i < bk {
			continue
		}

		v := majorityElement(c.nums)
		fmt.Println(c.nums, v)
	}
}
