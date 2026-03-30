type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    res := ""
    for _, s := range strs {
        l := strconv.Itoa(len(s))
        res += l + "#" + s
    }
    return res
}

func (s *Solution) Decode(encoded string) []string {
    res := make([]string, 0)
    i := 0
    for i < len(encoded){
        j := i 
        for encoded[j] != '#' {
            j++
        }
        l,_ := strconv.Atoi(encoded[i:j])
        i = j + 1
        j = i + l
        s := encoded[i:j]
        res = append(res, s)
        i += l
    }
    return res
}
