package main

import "fmt"

func singleNumber(nums []int) int {
	number := 0
	for i := 1; i < len(nums); i++ {
		if nums[number] == nums[i] {
			nums = append(nums[1:i], nums[i+1:]...)
			i = 0
		}
	}
	return nums[number]
}

func singleNumber1(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	nums := []int{1, 2, 1, 2, 4}
	fmt.Println(singleNumber1(nums))
}
