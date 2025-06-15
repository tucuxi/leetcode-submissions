func takeCharacters(s string, k int) int {
    var count, window [3]int
    
    for i := range s {
        count[s[i] - 'a']++
    }

    if min(count[0], count[1], count[2]) < k {
        return -1
    }

    res := len(s)

    i := 0
    for j := range s {
        window[s[j] - 'a']++
        for i <= j && min(count[0] - window[0], count[1] - window[1], count[2] - window[2]) < k {
            window[s[i] - 'a']--
            i++
        }
        res = min(res, len(s) - j + i - 1)
    }

    return res
}