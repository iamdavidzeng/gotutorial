package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]

		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			res := []int{nums[left], nums[right]}
			return res
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
