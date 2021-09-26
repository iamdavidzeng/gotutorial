package main

import "fmt"

func maxProfit(prices []int) int {
	minPriceIndex := 0
	maxPriceIndex := 0
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] >= prices[i+1] {
			if minPriceIndex < maxPriceIndex {
				profit += prices[maxPriceIndex] - prices[minPriceIndex]
			}
			minPriceIndex = i + 1
		} else {
			maxPriceIndex = i + 1
		}

		if maxPriceIndex == len(prices)-1 {
			profit += prices[maxPriceIndex] - prices[minPriceIndex]
		}
	}
	return profit
}

func main() {
	prices := []int{7, 6, 5, 4, 3, 1, 2}
	fmt.Println(maxProfit(prices))
}
