func topKFrequent(nums []int, k int) []int {
    h := make(map[int]int)
    for _,i := range nums {
        h[i]++
    }

    freq := make([][]int, len(nums) + 1)

    for v,i := range h {
        freq[i] = append(freq[i], v)
    }

    res := make([]int, 0)
    
    for i:= len(freq) - 1; i >=0; i-- {
        for _,v := range freq[i]{
            res = append(res, v)
            if len(res) == k {
                return res
            }
        }
    }

    return res
}
