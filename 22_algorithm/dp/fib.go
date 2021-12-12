package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

var memo = make(map[int]int)

func fibWithMemo(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fibWithMemo(n-1) + fibWithMemo(n-2)
	return memo[n]
}

func fibWithTable(n int) int {
	var dp = make(map[int]int)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fibWithPre(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	prev, curr := 1, 1
	for i := 3; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

func main() {
	fmt.Println(
		fibWithPre(20),
	)

}
