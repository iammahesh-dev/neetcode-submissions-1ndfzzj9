func groupAnagrams(strs []string) [][]string {
    res := make(map[[26]int][]string)
    for _, s := range strs {
        var count [26]int
        for _, c := range s{
            count[c-'a']++
        }
        res[count] = append(res[count], s)
    }
    result := make([][]string, 0)
    for _, g := range res {
        result = append(result, g)
    }

    return result
}
