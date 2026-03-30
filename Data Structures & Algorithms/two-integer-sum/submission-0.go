func twoSum(nums []int, target int) []int {
    h := make(map[int]int)
    for i, v := range nums {
        if j, ok := h[target-v]; ok {
            return []int{j, i}
        }
        h[v] = i
    }
    return []int{-1, -1}
}
