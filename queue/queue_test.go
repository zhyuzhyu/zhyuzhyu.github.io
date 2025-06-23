package queue

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
		{[]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4},
	}
	for _, c := range cases {
		fmt.Println(c.nums, c.k)
		fmt.Println(maxSlidingWindow(c.nums, c.k))
	}
}

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2},
	}
	for _, c := range cases {
		fmt.Println(c.nums, c.k)
		fmt.Println(topKFrequent(c.nums, c.k))
	}
}
