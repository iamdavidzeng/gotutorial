package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) < 0 {
		return []int{-1, -1}
	}
	result := []int{}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] == target:
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		result = append(result, -1)
	} else {
		result = append(result, left)
	}

	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] == target:
			left = mid + 1
		}
	}

	if right < 0 || nums[right] != target {
		result = append(result, -1)
	} else {
		result = append(result, right)
	}

	return result
}

func main() {

	nums := []int{5, 7, 7, 8, 8, 10}

	fmt.Println(searchRange(nums, 8))
}
