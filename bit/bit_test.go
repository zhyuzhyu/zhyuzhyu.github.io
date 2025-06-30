package bit

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		{[]int{2, 2, 1}},
		{[]int{4, 1, 2, 1, 2}},
		{[]int{1}},
	}
	for _, c := range cases {
		fmt.Println(c.nums)
		fmt.Println(singleNumber(c.nums))
	}
}
