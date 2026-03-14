func minOperations(s string) int {
    d := 0
    p := s[0]
    for i := range s {
        if s[i] != p {
            d++
        } 
        p ^= 1
    }
    return min(d, len(s) - d)
}