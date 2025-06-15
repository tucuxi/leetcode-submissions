func soupServings(n int) float64 {
    h := make(map[[2]int]float64)
    
    var dp func(int, int) float64
    dp = func(a, b int) float64 {
        if a <= 0 && b <= 0 {
            return 0.5
        }
        if a <= 0 {
            return 1.
        }
        if b <= 0 {
            return 0.
        }
        key := [2]int{a, b}
        if p, exists := h[key]; exists {
            return p
        }
        p := (dp(a - 4, b) + dp(a - 3, b - 1) + dp(a - 2, b - 2) + dp(a - 1, b - 3)) / 4.
        h[key] = p
        return p
    }
    
    m := (n + 24) / 25
    
    for i := 1; i <= m; i++ {
        if p := dp(i, i); p > 1. - .00001 {
            return p
        }
    }
    return dp(m, m)
}