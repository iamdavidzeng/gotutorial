package trans

import "math"

var Pi float64

func Init() float64 {
	Pi = 4 * math.Atan(1)
	return Pi
}
