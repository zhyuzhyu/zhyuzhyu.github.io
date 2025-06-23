package array

import "sort"

// https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func arraysMerge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums1[m-1] >= nums2[n-1] { //tail copy
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	if n > 0 {
		copy(nums1[:n], nums2[:n])
	}
}

// https://leetcode.cn/problems/remove-element/description/?envType=study-plan-v2&envId=top-interview-150
func removeElement(nums []int, val int) int {
	left := 0
	for i := range nums {
		if nums[i] != val {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/submissions/16825423/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	left := 0
	for i := range nums {
		if i+1 == len(nums) || nums[i] != nums[i+1] {
			nums[left] = nums[i]
			left++
		}
	}

	return left
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	left := 2
	for i := 2; i < n; i++ {
		if nums[i] != nums[left-2] {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

// https://leetcode.cn/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150
func majorityElement(nums []int) int {
	var candi, cnt int
	for _, v := range nums {
		if cnt == 0 {
			candi = v
			cnt = 1
		} else if v == candi {
			cnt++
		} else {
			cnt--
		}
	}
	return candi
}

// https://leetcode.cn/problems/move-zeroes/description/?envType=study-plan-v2&envId=top-100-liked
func moveZeroes(nums []int) {
	left := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[left] = nums[i]
			left++
		}
	}
	for i := left; i < len(nums); i++ {
		nums[i] = 0
	}
}

// https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-100-liked
func maxArea(height []int) int {
	var mArea, area int

	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			area = (j - i) * height[i]
			i++
		} else {
			area = (j - i) * height[j]
			j--
		}
		if area > mArea {
			mArea = area
		}
	}

	return mArea
}

// https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&envId=top-100-liked
func threeSum(nums []int) [][]int {
	rets := make([][]int, 0)
	sort.Ints(nums)
	if len(nums) < 3 {
		return rets
	}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			if j-1 > i && nums[j] == nums[j-1] {
				j++
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				rets = append(rets, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return rets
}

// https://leetcode.cn/problems/trapping-rain-water/description/?envType=study-plan-v2&envId=top-100-liked
func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	ret := 0
	for i := 0; i < n; i++ {
		ret += min(leftMax[i], rightMax[i]) - height[i]
	}

	return ret
}

// https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/description/
func numOfSubarrays(arr []int, k int, threshold int) int {
	// arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
	sum, total := 0, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i < k-1 {
			continue
		}
		if sum >= threshold*k {
			total++
		}
		sum -= arr[i-k+1]
	}
	return total
}

// https://leetcode.cn/problems/maximum-subarray/description/?envType=study-plan-v2&envId=top-100-liked
func maxSubArray(nums []int) int {
	// nums = [-2,1,-3,4,-1,2,1,-5,4]
	var arr sort.IntSlice

}
