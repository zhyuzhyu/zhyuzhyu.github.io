package strs

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-100-liked
// sliding window
func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]bool)
	n := len(s)
	ret := 0
	for i, j := 0, 0; i < n; i++ {
		if i != 0 {
			delete(mp, s[i-1])
		}
		for j < n && !mp[s[j]] {
			mp[s[j]] = true
			j++
		}
		if j-i > ret {
			ret = j - i
		}
	}
	return ret
}

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/?envType=study-plan-v2&envId=top-100-liked
func findAnagrams(s string, p string) []int {
	// s = "cbaebabacd", p = "abc"
	m, n := len(s), len(p)
	if m < n {
		return nil
	}

	var sCount, pCount [26]int
	for i := 0; i < n; i++ {
		sCount[s[i]-'a']++
		pCount[p[i]-'a']++
	}

	rets := make([]int, 0)
	if sCount == pCount { //array can compare directly
		rets = append(rets, 0)
	}
	for i := range s[:m-n] {
		sCount[s[i]-'a']--
		sCount[s[i+n]-'a']++
		if sCount == pCount {
			rets = append(rets, i+1)
		}
	}

	return rets
}

// https://leetcode.cn/problems/minimum-window-substring/description/?envType=study-plan-v2&envId=top-100-liked
func minWindow(s string, t string) string {
	// s = "ADOBECODEBANC", t = "ABC"
	mt, cnt := make(map[byte]int), make(map[byte]int)
	n := len(t)
	for _, v := range []byte(t) {
		mt[v]++
	}
	minLen := 0
	for i, j := 0, 0; j < len(s); {
		for k := i; k < j; k++ {
			cnt[s[k]]++
		}
	}
}
