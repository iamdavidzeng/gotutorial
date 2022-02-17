package main

import (
	"fmt"
	"sort"
)

func main() {
	lst := [][]int{
		[]int{5, 4},
		[]int{6, 4},
		[]int{6, 7},
		[]int{2, 3},
	}
	sort.Slice(lst, func(i, j int) bool {
		if lst[i][0] < lst[j][0] {
			return true
		} else if lst[i][0] == lst[j][0] {
			if lst[i][1] > lst[j][1] {
				return true
			}
		}
		return false
	})

	fmt.Println(lst)
}
