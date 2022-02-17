package main

import (
	"fmt"
	"math"
)

func longestCommonSubsequence(s1, s2 string) int {

	var dp func(s1, s2 string, i, j int) int
	dp = func(s1, s2 string, i, j int) int {
		if i == len(s1) || j == len(s2) {
			return 0
		}

		if string(s1[i]) == string(s2[j]) {
			return 1 + dp(s1, s2, i+1, j+1)
		} else {
			return max(
				dp(s1, s2, i+1, j),
				dp(s1, s2, i, j+1),
			)
		}
	}

	return dp(s1, s2, 0, 0)
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
