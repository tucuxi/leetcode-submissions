func countKConstraintSubstrings(s string, k int) int {
    res := 0
    zeros := 0
    ones := 0
    p := 0
    for q := 0; q < len(s); q++ {
        if s[q] == '0' {
            zeros++
        } else {
            ones++
        }
        for ; zeros > k && ones > k; p++ {
            if s[p] == '0' {
                zeros--
            } else {
                ones--
            }
        }
        res += q - p + 1
    }
    return res
}