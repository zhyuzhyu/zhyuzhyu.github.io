package main

import (
	"maps"
	"strings"
)

// pattern = "abba", s = "dog cat cat dog"
func wordPattern(pattern string, s string) bool {
	arrs := strings.Split(s, " ")
	if len([]byte(pattern)) != len(arrs) {
		return false
	}
}

func main() {

}
