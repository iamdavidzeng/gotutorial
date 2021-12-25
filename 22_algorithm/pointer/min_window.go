package main

import (
	"fmt"
	"math"
)

func minWindow(s, t string) string {
	need, window := make(map[string]int), make(map[string]int)
	for _, v := range t {
		need[string(v)]++
	}
	valid := 0
	left, right := 0, 0
	start, length := 0, math.MaxInt32
	for right < len(s) {
		c := string(s[right])
		right++

		if need[c] == 1 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			d := string(s[left])
			left++
			if need[d] == 1 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start:(start + length)]
}

func main() {
	fmt.Println(minWindow("EBBANCF", "ABC"))
}
