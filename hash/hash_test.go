package hash

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9},
		{nums: []int{3, 2, 4}, target: 6},
	}
	for _, c := range cases {
		fmt.Println(c.nums, c.target)
		fmt.Println(twoSum(c.nums, c.target))
	}
}

func TestGroupAnagrams(t *testing.T) {
	cases := []struct {
		strs []string
	}{
		{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
	}
	for _, c := range cases {
		fmt.Println(c.strs)
		fmt.Println(groupAnagrams(c.strs))
	}
}

func TestLongestConsecutive(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		{nums: []int{100, 4, 200, 1, 3, 2}},
		{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
		{nums: []int{1, 3, 5, 2, 4}},
		{nums: []int{-8, -4, 9, 9, 4, 6, 1, -4, -1, 6, 8}},
	}
	for _, c := range cases {
		fmt.Println(c.nums)
		fmt.Println(longestConsecutive(c.nums))
	}
}

func TestSubarraySum(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
	}{
		//{nums: []int{1, 1, 1}, k: 2},
		//{nums: []int{1, 2, 3}, k: 3},
		{nums: []int{1, -1, 0}, k: 0},
	}
	for _, c := range cases {
		fmt.Println(c.nums, c.k)
		fmt.Println(subarraySum(c.nums, c.k))
	}
}
