func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	longest := 0
	for _,v := range nums {
		m[v] = true
	}

	for _,v := range nums {
		if _, ok := m[v-1]; !ok {
			length := 0
			for m[v+length] {
				length++
			}
			longest = max(length, longest)
		}
	}
	return longest
}
