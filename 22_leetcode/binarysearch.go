package main

import "fmt"

/*
使用缩小搜索空间的方式来解决搜索：

首先确定使用闭合的搜索空间[left, right]，
所以：
1. right的取值为len - 1
2. for循环的条件为left <= right(因为left只有大于right的时候才可以认为搜索空间为空)
3. 当mid != target的情况，left、right都是加一或者减一

当找到mid之后，将mid的值进行返回，如果整个搜索空间内都没找到target，则返回-1.
*/
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}

	fmt.Println(search(nums, 9))
}
