package main

import "fmt"

func maxProfit(prices []int) int {
	min_price_index := 0
	max_price_index := 0
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] >= prices[i+1] {
			if min_price_index < max_price_index {
				profit += prices[max_price_index] - prices[min_price_index]
			}
			min_price_index = i + 1
		} else {
			max_price_index = i + 1
		}

		if max_price_index == len(prices)-1 {
			profit += prices[max_price_index] - prices[min_price_index]
		}
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
