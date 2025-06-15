func numSmallerByFrequency(queries []string, words []string) []int {
    fw := mapStrings(words, f)
    sort.Ints(fw)
    return mapStrings(queries, func(q string) int {
        return len(fw) - sort.SearchInts(fw, f(q)+1)
    })
}

func mapStrings(s []string, f func(s string) int) []int {
    res := make([]int, len(s))
    for i := range s {
        res[i] = f(s[i])
    }
    return res
}

func f(s string) int {
    smc := s[0]
    res := 0
 
    for i := range s {
        if s[i] < smc {
            res = 1
            smc = s[i]
        } else if s[i] == smc {
            res++
        }
    }
    return res
}