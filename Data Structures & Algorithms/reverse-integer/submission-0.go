func reverse(x int) int {
	sign := 1 
	if x < 0 {
		sign = -1
	}

	x = abs(x)

	res:=0 
	for x > 0 {
		res = res * 10 + x % 10
		x /= 10
	}
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	return sign * res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
