package main

import "fmt"

type RunCase struct {
	nums1, nums2 []int
	m, n         int
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums2[n-1] >= nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}

	if n > 0 {
		copy(nums1[:n], nums2[:n])
	}
}

func main() {
	var cases []RunCase
	cases = append(cases, RunCase{
		nums1: []int{1, 2, 3, 0, 0, 0},
		nums2: []int{2, 5, 6},
		m:     3,
		n:     3,
	})

	cases = append(cases, RunCase{
		nums1: []int{1},
		nums2: []int{},
		m:     1,
		n:     0,
	})

	cases = append(cases, RunCase{
		nums1: []int{0},
		nums2: []int{1},
		m:     0,
		n:     1,
	})

	cases = append(cases, RunCase{
		nums1: []int{2, 0},
		nums2: []int{1},
		m:     1,
		n:     1,
	})

	bk := 0
	for i, c := range cases {
		if i < bk {
			continue
		}

		merge(c.nums1, c.m, c.nums2, c.n)
		fmt.Println(c.nums1)
	}
}
