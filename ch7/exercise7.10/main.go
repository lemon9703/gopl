package main

import (
	"fmt"
	"sort"
)

type str []byte

func (s str) Len() int           { return len(s) }
func (s str) Less(i, j int) bool { return s[i] < s[j] }
func (s str) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	a := str("abccba")
	fmt.Println(IsPalindrome(a))

	b := str("abcdefg")
	fmt.Println(IsPalindrome(b))
}
