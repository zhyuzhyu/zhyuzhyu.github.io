package main

import "fmt"

func rotate(nums []int, k int) {
	l := len(nums)
	j := l - k%l
	if j == l {
		return
	}

	arr := make([]int, 0, l)
	arr = append(arr, nums[j:]...)
	arr = append(arr, nums[:j]...)
	copy(nums, arr)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}
