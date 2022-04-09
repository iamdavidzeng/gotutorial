package main

import "fmt"

func findAnagrams(s string, p string) []int {
	window, need := map[byte]int{}, map[byte]int{}
	for i := range p {
		need[p[i]]++
	}

	res := []int{}
	valid := 0

	p1, p2 := 0, 0
	for p1 < len(s) {
		s1 := s[p1]
		if _, ok := need[s1]; ok {
			window[s1]++
			if window[s1] == need[s1] {
				valid++
			}
		}
		p1++

		for (p1 - p2) >= len(p) {
			if valid == len(need) {
				res = append(res, p2)
			}
			s2 := s[p2]
			if _, ok := need[s2]; ok {
				if window[s2] == need[s2] {
					valid--
				}
				window[s2]--
			}
			p2++
		}
	}
	return res
}

func main() {
	fmt.Println(findAnagrams("baa", "aa"))
}
