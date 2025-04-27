func commonFactors(a int, b int) int {
    res := 0
    limit := min(a, b)
    for i := 1; i <= limit; i++ {
        if a % i == 0 && b % i == 0 {
            res++
        }
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}