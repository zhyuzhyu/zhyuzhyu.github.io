package array

import (
	"fmt"
	"testing"
)

func TestArraysMerge(t *testing.T) {
	cases := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}{
		{[]int{1, 2, 3, 6, 0, 0, 0}, 4, []int{2, 5, 6}, 3},
		{[]int{1}, 1, []int{}, 0},
		{[]int{0}, 0, []int{1}, 1},
		{[]int{2, 0}, 1, []int{1}, 1},
	}
	for _, c := range cases {
		fmt.Println("prev: ", c.nums1, c.m, c.nums2, c.n)
		arraysMerge(c.nums1, c.m, c.nums2, c.n)
		fmt.Println("post: ", c.nums1)
	}
}

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		nums []int
		val  int
	}{
		{[]int{3, 2, 2, 3}, 3},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
	}
	for _, c := range cases {
		fmt.Println(c)
		fmt.Println(c.nums[:removeElement(c.nums, c.val)])
	}
}

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		{[]int{1, 1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
	}
	for _, c := range cases {
		fmt.Println(c.nums)
		fmt.Println(c.nums[:removeDuplicates(c.nums)])
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
	}
	for _, c := range cases {
		fmt.Println(c.nums)
		fmt.Println(c.nums[:removeDuplicates2(c.nums)])
	}
}

func TestMajorityElement(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		{[]int{3, 2, 3}},
		{[]int{2, 2, 1, 1, 1, 2, 2}},
		{[]int{6, 5, 5}},
	}
	for _, c := range cases {
		fmt.Println(c.nums)
		fmt.Println(majorityElement(c.nums))
	}
}

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		{[]int{0, 1, 0, 3, 12}},
	}
	for _, c := range cases {
		fmt.Println(c.nums)
		moveZeroes(c.nums)
		fmt.Println(c.nums)
	}
}

func TestMaxArea(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
	}
	for _, c := range cases {
		fmt.Println(c.nums)
		fmt.Println(maxArea(c.nums))
	}
}

func TestThreeSum(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		//{[]int{-1, 0, 1, 2, -1, -4}},
		{[]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}},
	}
	for _, c := range cases {
		fmt.Println(c.nums)
		fmt.Println(threeSum(c.nums))
	}
}

func TestTrap(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
		{[]int{4, 2, 0, 3, 2, 5}},
	}
	for _, c := range cases {
		fmt.Println(c.nums)
		fmt.Println(trap(c.nums))
	}
}

func TestNumOfSubarrays(t *testing.T) {
	cases := []struct {
		nums         []int
		k, threshold int
	}{
		{[]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4},
	}
	for _, c := range cases {
		fmt.Println(c.nums, c.k, c.threshold)
		fmt.Println(numOfSubarrays(c.nums, c.k, c.threshold))
	}
}
