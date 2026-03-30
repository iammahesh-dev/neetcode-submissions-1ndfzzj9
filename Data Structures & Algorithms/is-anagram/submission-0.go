func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    chS, chT := make(map[byte]int), make(map[byte]int)
    for i := 0; i < len(s); i++ {
        chS[s[i]]++
        chT[t[i]]++
    }

    for i := 0; i < len(s); i++ {
        if chS[s[i]] != chT[s[i]] {
            return false
        }
    }
    return true
}
