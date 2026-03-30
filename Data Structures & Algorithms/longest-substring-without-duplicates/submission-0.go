func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	l, res := 0, 0
	for r:= 0; r < len(s); r++{
		if idx, ok := m[s[r]]; ok {
			l = max(idx + 1, l)
		}
		m[s[r]] = r
		res = max(r - l + 1, res)
	}
	return res
}
