func maximumLength(s string) int {
    h := make(map[string]int)
    var b []byte

    for i := range s {
        if len(b) > 0 && s[i] != b[0] {
            b = nil
        }
        b = append(b, s[i])
        for j := 1; j <= len(b); j++ {
            h[string(b[:j])]++
        }
    }
    
    maxLen := -1
    for subStr, count := range h {
        if count >= 3 {
            maxLen = max(maxLen, len(subStr))
        }
    }

    return maxLen
}