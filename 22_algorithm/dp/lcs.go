package main

import (
	"fmt"
	"math"
)

// 自顶向下
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

// 自底向上
func lcs(s1, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(
					dp[i-1][j],
					dp[i][j-1],
				)
			}
		}
	}

	return dp[m][n]
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
