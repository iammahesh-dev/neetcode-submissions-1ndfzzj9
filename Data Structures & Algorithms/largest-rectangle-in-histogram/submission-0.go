func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := []int{}
	n := len(heights)
	for i:=0; i <= n; i++ {
		for len(stack) > 0 && ( i == n || heights[stack[len(stack) - 1]] >= heights[i]){
			ht := heights[stack[len(stack) - 1]]
			stack = stack[:len(stack) - 1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack) - 1] - 1
			}
			area := ht * width
			maxArea = max(area, maxArea)
		}
		if i < n {
			stack = append(stack, i)
		}
	}
	return maxArea
}
