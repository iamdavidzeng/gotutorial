package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	fmt.Println(intervals)

	counter := 0
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			counter += 1
		} else {
			end = intervals[i][1]
		}
	}
	return counter
}

func main() {
	intervals := [][]int{[]int{1, 100}, []int{11, 22}, []int{1, 11}, []int{2, 12}}
	fmt.Println(eraseOverlapIntervals(intervals))
}
