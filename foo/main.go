package main

import "fmt"

func findAnagrams(s string, p string) []int {
	window, need := map[byte]int{}, map[byte]int{}
	for i := range p {
		need[p[i]]++
	}

	res := []int{}
	left, right := 0, 0
	valid := 0
	for right < len(s) {
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		for (right - left) >= len(p) {
			if valid == len(need) {
				fmt.Println(valid, need, window)
				res = append(res, left)
			}

			d := s[left]
			left++

			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findAnagrams("baa", "aa"))
}
