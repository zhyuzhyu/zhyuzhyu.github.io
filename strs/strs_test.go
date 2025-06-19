package strs

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []string{
		//"abcabcbb",
		"dvdf",
	}
	for _, c := range cases {
		fmt.Println(c)
		fmt.Println(lengthOfLongestSubstring(c))
	}
}

func TestFindAnagrams(t *testing.T) {
	cases := []struct {
		s, p string
	}{
		{"cbaebabacd", "abc"},
		{"abab", "ab"},
	}
	for _, c := range cases {
		fmt.Println(c.s, c.p)
		fmt.Println(findAnagrams(c.s, c.p))
	}
}
