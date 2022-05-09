package tDirectory

func Sum(set ...int) int {
	var s int
	for _, v := range set {
		s += v
	}
	return s
}

func Mul(set ...int) int {
	var s = 1
	for _, v := range set {
		s *= v
	}
	return s
}
