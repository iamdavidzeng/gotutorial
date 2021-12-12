package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {

	memo := make(map[int]int, len(coins)+1)

	return change(memo, coins, amount)
}

func change(memo map[int]int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if value, ok := memo[amount]; ok {
		return value
	}
	counter := 65535
	for _, coin := range coins {
		result := change(memo, coins, amount-coin)
		if result == -1 {
			continue
		}
		counter = min(counter, result+1)
	}
	if counter == 65535 {
		memo[amount] = -1
		return -1
	}
	memo[amount] = counter
	return counter
}

func coinChangeWithDP(coins []int, amount int) int {
	if amount < 1 && len(coins) < 1 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, c := range coins {
			if i >= c && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(coinChangeWithDP([]int{1, 2, 5}, 11))
}
