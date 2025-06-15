func numberOfSubstrings(s string, k int) int {
    var count [26]int
    maxCount := 0
    res := 0
    left := 0
    for right := range s {
        chr := s[right] - 'a'
        count[chr]++
        maxCount = max(maxCount, count[chr])
        for maxCount >= k {
            res += len(s) - right
            chl := s[left] - 'a'
            count[chl]--
            maxCount = 0
            for _, c := range count {
                maxCount = max(maxCount, c)
            }
            left++
        }
    }
    return res
}