func findTheLongestBalancedSubstring(s string) int {
    count := 0
    zeros, ones := 0, 0
    for i := range s {
        if s[i] == '0' {
            if ones > 0 {
                count = max(count, 2 * min(zeros, ones))
                zeros, ones = 0, 0
            }
            zeros++
        } else {
            ones++
        }
    }
    return max(count, 2 * min(zeros, ones))
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}