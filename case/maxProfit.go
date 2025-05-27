package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	start, profit := prices[0], 0

	for _, v := range prices[1:] {
		delta := v - start
		if delta > profit {
			profit = delta
		} else if delta < 0 {
			start = v
		}
	}

	return profit
}

func main() {
	// [7,1,5,3,6,4]
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
