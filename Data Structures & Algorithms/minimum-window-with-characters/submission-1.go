func minWindow(s string, t string) string {
    countT := make(map[byte]int)
	for i:=0; i < len(t); i++ {
		countT[t[i]]++
	}
	res, minLen := "", math.MaxInt32
	have, need := 0, len(countT)
	winF := make(map[byte]int)
	l:=0
	for r:=0; r< len(s); r++{
		winF[s[r]]++
		if countT[s[r]] > 0 && winF[s[r]] == countT[s[r]]{
			have++
		}
		for have == need {
			if r - l + 1 < minLen {
				res = s[l: r + 1]
				minLen = r - l + 1
			}

			winF[s[l]]--
			if countT[s[l]] > 0 && winF[s[l]] < countT[s[l]]{
				have--
			}
			l++
		}
	}
	return res
}
