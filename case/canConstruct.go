package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	mp := make(map[byte]int)
	for _, c := range []byte(magazine) {
		mp[c]++
	}
	for _, c := range []byte(ransomNote) {
		mp[c]--
		if mp[c] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
