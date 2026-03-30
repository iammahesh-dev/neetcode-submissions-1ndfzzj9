func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1 / x
        n = -n
    }
	if n == 0 {
        return 1
    }
    
    if n == 1 {
        return x
    }
	result := 1.0
	for n > 0 {
		if n % 2 != 0 {
			result *= x
		}
		x *= x
		n /= 2
	}
	return result
}
