func characterReplacement(s string, k int) int {
    var count [26]int
    left := 0
    res := 0
    for right := range s {
        count[s[right] - 'A']++
        for {
            maxCount := 0
            for _, c := range count {
                maxCount = max(c, maxCount)
            }
            if right - left - maxCount < k {
                break
            }
            count[s[left] - 'A']--
            left++
        }
        res = max(right - left + 1, res)
    }
    return res
}