package algorithm

func IntReverse(target int) int {
	y := 0
	for target != 0 {
		y = y*10 + target%10
		if !(-(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		target /= 10
	}
	return y
}
