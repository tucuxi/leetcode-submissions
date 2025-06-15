func beautySum(s string) int {
    prefix := make([][26]int, len(s) + 1)
    for i := range s {
        prefix[i + 1] = prefix[i]
        prefix[i + 1][s[i] - 'a']++
    }
    res := 0
    for i := range prefix {
        for j := range i {
            var subStringFreqs [26]int
            for k := range subStringFreqs {
                subStringFreqs[k] = prefix[i][k] - prefix[j][k]
            }
            minFreq := math.MaxInt
            maxFreq := 0
            for _, f := range subStringFreqs {
                if f > 0 {
                    minFreq = min(minFreq, f)
                    maxFreq = max(maxFreq, f)
                }
            }
            if minFreq < maxFreq {
                res += maxFreq - minFreq
            }
        }
    }
    return res
}