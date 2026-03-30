func hasDuplicate(nums []int) bool {
    h := make(map[int]bool)
    for _, v := range nums {
        if _, ok := h[v]; ok {
            return true
        }
        h[v] = true
    }
    return false
}
