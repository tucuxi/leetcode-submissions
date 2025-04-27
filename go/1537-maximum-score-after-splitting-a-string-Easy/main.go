func maxScore(s string) int {
    res := 0
    zeroes := prefixSum(s)
    ones := 0
    for i := len(s) - 1; i > 0; i-- {
        if s[i] == '1' {
            ones++
        }
        res = max(res, ones + zeroes[i])
    }
    return res
}

func prefixSum(s string) []int {
    res := make([]int, len(s))
    zeroes := 0
    for i := range s {
        res[i] = zeroes
        if s[i] == '0' {
            zeroes++
        }
    }
    return res
}