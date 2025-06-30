package bit

/*
x ^ x = 0
x ^ 0 = x
*/

// https://leetcode.cn/problems/single-number/description/?envType=study-plan-v2&envId=top-100-liked
func singleNumber(nums []int) int {
	product := 0
	for i := 0; i < len(nums); i++ {
		product ^= nums[i]
	}
	return product
}
