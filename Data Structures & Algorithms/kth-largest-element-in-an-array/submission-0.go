func minHeap(h []int, val int) []int{
    h = append(h, val)
    index := len(h) - 1
    for index > 0 && h[(index - 1) / 2] > h[index]{
        h[(index - 1) / 2], h[index] = h[index], h[(index  - 1) / 2]
        index = (index - 1) / 2
    }
    return h
}

func delTop(h []int) []int{
    h[0] = h[len(h) - 1]
    h = h[:len(h) - 1]
    index := 0
    for {
        smallest := index
        left := 2 * index + 1
        right := 2 * index + 2
        if left < len(h) && h[left] < h[smallest]{
            smallest = left
        }
        if right < len(h) && h[right] < h[smallest]{
            smallest = right
        }
        if smallest == index {
            break
        }

        h[smallest], h[index] = h[index], h[smallest]
        index = smallest
    }
    return h
}

func findKthLargest(nums []int, k int) int {
    h := make([]int, 0)
    for i:=0;i<len(nums);i++ {
        h = minHeap(h, nums[i])
        if len(h) > k {
            h = delTop(h)
        }
    }
    return h[0]
}