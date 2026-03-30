func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0

	for _, p := range piles {
		if p > r {
			r = p
		}
	}

	res:= r

	for l <= r {
		k := l + (r - l) / 2
		hours := 0
		for _,p := range piles {
			hours += int(math.Ceil(float64(p) / float64(k)))
		}
		if hours <= h {
			res = k
			r = k - 1

		} else {
			l = k + 1
		}
	}

	return res

}
