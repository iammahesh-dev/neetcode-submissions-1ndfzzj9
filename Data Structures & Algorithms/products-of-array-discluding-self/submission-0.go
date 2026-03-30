func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    pre := 1
    for i,v := range nums {
        res[i] = pre
        pre *= v
    }

    post := 1
    for i:=len(nums) - 1; i >= 0 ; i-- {
        res[i] *= post
        post *= nums[i]
    }

    return res
}
