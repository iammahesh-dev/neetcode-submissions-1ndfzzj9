func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int, 0)
	q := make([]int, 0)
	l:=0
	for r,v := range nums {
		for len(q) > 0 && nums[q[len(q) - 1]] < v {
			q = q[:len(q) - 1]
		}

		q = append(q, r)
		if q[0] < l {
			q = q[1:]
		}

		if r + 1 >= k {
			res = append(res, nums[q[0]])
			l++
		}
		r++
	}
	return res
}
