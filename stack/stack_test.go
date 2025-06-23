package stack

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		str string
	}{
		{"([])"},
		{"(]"},
		{"){"},
		{"(){}}{"},
		{"()"},
	}
	for _, c := range cases {
		fmt.Println(c.str)
		fmt.Println(isValid(c.str))
	}
}

func TestCalculate(t *testing.T) {
	cases := []struct {
		str string
	}{
		{"3+2*2"},
	}
	for _, c := range cases {
		fmt.Println(c.str)
		fmt.Println(calculate(c.str))
	}
}

func TestNextGreaterElement(t *testing.T) {
	cases := []struct {
		nums1, nums2 []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}},
	}
	for _, c := range cases {
		fmt.Println(c.nums1, c.nums2)
		fmt.Println(nextGreaterElement(c.nums1, c.nums2))
	}
}

func TestDailyTemperatures(t *testing.T) {
	cases := []struct {
		temperatures []int
	}{
		//{[]int{73, 74, 75, 71, 69, 72, 76, 73}},
		{[]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}},
	}
	for _, c := range cases {
		fmt.Println(c.temperatures)
		fmt.Println(dailyTemperatures(c.temperatures))
	}
}
