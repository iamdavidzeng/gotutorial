package algorithm

func MergeSort(lst []int) []int {
	length := len(lst)

	if length <= 1 {
		return lst
	} else {
		pivot := length / 2
		return merge(MergeSort(lst[:pivot]), MergeSort(lst[pivot:]))
	}
}

func merge(left, right []int) []int {
	new := []int{}

	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			new = append(new, left[0])
			left = left[1:]
		} else {
			new = append(new, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		new = append(new, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		new = append(new, right[0])
		right = right[1:]
	}

	return new
}
