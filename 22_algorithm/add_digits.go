package algorithm

func AddDigits(num int) int {
	if num < 10 {
		return num
	} else {
		mid := 0
		for num != 0 {
			mid += num % 10
			num /= 10
		}
		return AddDigits(mid)
	}
}
