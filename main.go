package quote

func Sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
