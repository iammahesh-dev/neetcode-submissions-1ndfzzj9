func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2){
		return false
	}
	counts1, counts2 := make([]int, 26), make([]int, 26)
	matches := 0
	for i:=0; i < len(s1); i++{
		counts1[s1[i] - 'a']++
		counts2[s2[i] - 'a']++
	}

	for i:=0; i < 26; i++ {
		if counts1[i] == counts2[i]{
			matches++
		}
	}

	l:=0
	for r:=len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}
		idx := s2[r] - 'a'
		counts2[idx]++
		if counts1[idx] == counts2[idx] {
			matches++
		} else if counts1[idx] + 1 == counts2[idx]{
			matches--
		}

		idx = s2[l] - 'a'
		counts2[idx]--
		if counts1[idx] == counts2[idx] {
			matches++
		} else if counts1[idx] - 1 == counts2[idx]{
			matches--
		}
		l++
	}
	return matches == 26
}
