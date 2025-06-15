func distributeCandies(n int, limit int) int {
    res := 0
    for i := min(n, limit); i >= 0; i-- {
        for j := min(n - i, limit); j >= 0; j-- {
            if n - i - j <= limit {
                res++
            }   
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