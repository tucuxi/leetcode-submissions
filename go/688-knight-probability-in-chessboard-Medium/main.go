func knightProbability(n int, k int, row int, column int) float64 {
    steps := [][]int{{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}}
    p := make2D(n)
    p[row][column] = 1.
    pp := make2D(n)
    for i := 0; i < k; i++ {
        for r := range p {
            copy(pp[r], p[r])
            fill(p[r], 0)
        }
        for r := range p {
            for c := range p[r] {
                for _, step := range steps {
                    rr, cc := r - step[0], c - step[1]
                    if rr >= 0 && rr < n && cc >= 0 && cc < n {
                        p[r][c] += pp[rr][cc] / 8
                    }
                }
            }
        }
    }
    return sum(p)
}

func make2D(n int) [][]float64 {
    res := make([][]float64, n)
    for i := range res {
        res[i] = make([]float64, n)
    }
    return res
}

func fill(a []float64, v float64) {
    for i := range a {
        a[i] = v
    }
}

func sum(a [][]float64) float64 {
    res := 0.
    for _, r := range a {
        for _, c := range r {
            res += c
        }
    }
    return res
}