package main

import (
	"fmt"
)

func palindrome(s string, l, r int) string {
	for l > 0 && r < len(s) && s[l-1] == s[r+1] {
		l--
		r++
	}
	return s[l : r+1]
}

func isPalindrome1(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "cbbc"
	fmt.Println(isPalindrome1(s))
}
