func countSubstrings(s string) int {
    res := 0

    for i := 0; i < len(s); i++ {
        // Odd-length
        l, r := i, i
        for l >= 0 && r < len(s) && s[l] == s[r] {
            res++
            l--
            r++
        }

        // Even-length
        l, r = i, i+1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            res++
            l--
            r++
        }
    }

    return res
}
