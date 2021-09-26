package algorithm

func IsHappy(n int) bool {
	history := make(map[int]bool)

	for n != 1 {
		if _, ok := history[n]; ok {
			return false
		}
		history[n] = true

		sumValue := 0
		for n != 0 {
			sumValue += (n % 10) * (n % 10)
			n /= 10
		}

		n = sumValue
	}
	return true
}
