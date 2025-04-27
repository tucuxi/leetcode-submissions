// n <= 10

func numberOfSets(n int, maxDistance int, roads [][]int) int {
    check := func(mask int) bool {
        w := make([][]int, n)
        for i := 0; i < n; i++ {
            w[i] = make([]int, n)
            for j := 0; j < n; j++ {
                if j != i {
                    w[i][j] = 10_000
                }
            }
        }
        for _, r := range roads {
            i, j := r[0], r[1]
            w[i][j] = min(w[i][j], r[2])
            w[j][i] = min(w[j][i], r[2])
        }
        for k := 0; k < n; k++ {
            if mask & (1 << k) != 0 {
                for i := 0; i < n; i++ {
                    if mask & (1 << i) != 0 {
                        for j := 0; j < n; j++ {
                            if mask & (1 << j) != 0 {
                                w[i][j] = min(w[i][j], w[i][k] + w[k][j])
                            }
                        }
                    }
                }
            }
        }
        for i := 0; i < n; i++ {
            if mask & (1 << i) != 0 {
                for j := i + 1; j < n; j++ {
                    if mask & (1 << j) != 0 {
                        if w[i][j] > maxDistance {
                            return false
                        }
                    }
                }
            }
        }
        return true
    }
    
    res := 0
    for mask := 1 << n - 1; mask >= 0; mask-- {
        if check(mask) {
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