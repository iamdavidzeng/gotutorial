package main

import (
	"fmt"
)

func shipWithinDays(weights []int, days int) int {
	left, right := 1, getSum(weights)

	for left < right {
		mid := left + (right-left)/2
		if getDays(weights, mid, days) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func getDays(weights []int, speed, days int) bool {
	counter := 1
	cap := 0
	for i := range weights {
		if weights[i] > speed {
			return false
		}
		if cap+weights[i] <= speed {
			cap += weights[i]
		} else {
			cap = weights[i]
			days += 1
		}
	}
	return counter <= days
}

func getSum(nums []int) int {
	value := 0
	for _, v := range nums {
		value += v
	}
	return value
}

func main() {
	nums := []int{1, 2, 3, 1, 1}
	fmt.Println(shipWithinDays(nums, 4))
}
