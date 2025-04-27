func camelMatch(queries []string, pattern string) []bool {
    var matches []bool
    
    for _, q := range queries {
        i := 0
        j := 0
        for i < len(q) && j < len(pattern) {
            if q[i] == pattern[j] {
                i++
                j++
            } else if q[i] >= 'a' {
                i++
            } else {
                break
            }
        }
        for i < len(q) && q[i] >= 'a' {
            i++
        }
        matches = append(matches, i == len(q) && j == len(pattern))
    }
    
    return matches
}