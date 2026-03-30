func maxSubArray(nums []int) int {
    maxSub, curSum := nums[0], 0
	for _, v := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += v
		maxSub = max(curSum, maxSub)
	}
	return maxSub
}
