func champagneTower(poured int, query_row int, query_glass int) float64 {
    current := []float64{float64(poured)}
    for r := range query_row {
        next := make([]float64, r + 2)
        for c := range r + 1 {
            if q := (current[c] - 1) / 2; q > 0 {
                next[c] += q
                next[c + 1] += q
            }
        }
        current = next
    }
    return min(1.0, current[query_glass])
}