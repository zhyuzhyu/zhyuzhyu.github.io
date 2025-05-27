package main

import "fmt"

// [3,2,1,0,4]
func canJump(nums []int) bool {
	l := len(nums)
	for i, cnt := 0, 0; cnt < l; {
		if nums[i] == 0 && i != l-1 {
			return false
		} else if i+nums[i] >= l {
			return true
		}
		i += nums[i]
		cnt++
	}

	return false
}

func main() {
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
