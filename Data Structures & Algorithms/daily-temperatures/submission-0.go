func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(temperatures))
	for i,v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack) - 1]]{
			stackInd := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			res[stackInd] = i - stackInd
		}
		stack = append(stack, i)
	}
	return res
}
