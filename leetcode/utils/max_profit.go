package main

import "fmt"

func maxProfit(prices []int) int {
	min_price_index := 0
	max_price_index := 0
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] >= prices[i+1] {
			min_price_index = i + 1
		} else {
			max_price_index = i + 1
		}

		if min_price_index == len(prices)-1 {
			return 0
		} else if max_price_index == len(prices)-1 {
			return prices[max_price_index] - prices[min_price_index]
		}
	}

	return profit
}

func main() {
	prices := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))
}
