func average(salary []int) float64 {
    min, max, total := math.MaxInt, math.MinInt, 0
    for _, s := range salary {
        if s < min {
            min = s
        }
        if s > max {
            max = s
        }
        total += s
    }
    return float64(total - min - max) / float64(len(salary) - 2)
}