func champagneTower(poured int, query_row int, query_glass int) float64 {
    champagne := make([][]float64, query_row + 1)
    champagne[0] = make([]float64, 1)
    champagne[0][0] = float64(poured)
    for r := 0; r < query_row; r++ {
        champagne[r + 1] = make([]float64, r + 2)
        for c := 0; c <= r; c++ {
            if q := (champagne[r][c] - 1) / 2; q > 0 {
                champagne[r + 1][c] += q
                champagne[r + 1][c + 1] += q
            }
        }
    }
    return math.Min(1, champagne[query_row][query_glass])
}