package main

func minEatingSpeed(piles []int, h int) int {
	left, right := 0, getMax(piles)

	for left < right {
		mid := left + (right-left)/2
		if canFinish(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canFinish(nums []int, speed, hour int) bool {
	count := 0
	for i := 0; i < hour; i++ {
		if speed >= nums[i] {
			count += 1
		}
	}
	return count == len(nums)
}

func getMax(nums []int) int {
	value := 0
	for _, v := range nums {
		if v >= value {
			value = v
		}
	}
	return value
}

// func main() {
// 	nums := []int{30, 11, 23, 4, 20}

// 	fmt.Println(minEatingSpeed(nums, 5))
// }
