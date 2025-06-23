package hash

import (
	"sort"
)

// https://leetcode.cn/problems/two-sum/description/?envType=study-plan-v2&envId=top-100-liked
func twoSum(nums []int, target int) []int {
	nums_set := map[int]int{}
	for i, v := range nums {
		if j, ok := nums_set[target-v]; ok {
			return []int{j, i}
		}
		nums_set[v] = i
	}
	return nil
}

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
func groupAnagrams(strs []string) [][]string {
	strs_set := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sorted_str := string(s)
		if _, ok := strs_set[sorted_str]; !ok {
			group := []string{str}
			strs_set[sorted_str] = group
		} else {
			group := strs_set[sorted_str]
			group = append(group, str)
			strs_set[sorted_str] = group
		}
	}
	rets := make([][]string, 0, len(strs_set))
	for _, v := range strs_set {
		rets = append(rets, v)
	}
	return rets
}

// https://leetcode.cn/problems/longest-consecutive-sequence/?envType=study-plan-v2&envId=top-100-liked
func longestConsecutive(nums []int) int {
	nums_set := make(map[int]int)
	for _, v := range nums {
		nums_set[v] = 1
	}

	cnt := 0
	for k := range nums_set {
		if _, ok := nums_set[k-1]; ok {
			continue
		}
		for j := k + 1; ; j++ {
			if _, ok := nums_set[j]; ok {
				nums_set[k]++
				continue
			}
			break
		}
		if nums_set[k] > cnt {
			cnt = nums_set[k]
		}
	}

	return cnt
}

// https://leetcode.cn/problems/subarray-sum-equals-k/description/?envType=study-plan-v2&envId=top-100-liked
func subarraySum(nums []int, k int) int {
	// nums = [1,-1,0], k = 0
	pre := 0
	mPre := make(map[int]int)
	total := 0
	mPre[0] = 1
	for _, v := range nums {
		pre += v
		if v2, ok := mPre[pre-k]; ok {
			total += v2
		}
		mPre[pre]++
	}
	return total
}
