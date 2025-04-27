func findRepeatedDnaSequences(s string) []string {
    frequencies := map[string]int{}
    const length = 10
    for i := length; i <= len(s); i++ {
        frequencies[s[i - length:i]]++
    }
    res := []string{}
    for sequence, count := range frequencies {
        if count > 1 {
            res = append(res, sequence)
        }
    }
    return res
}