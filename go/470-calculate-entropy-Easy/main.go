func calculateEntropy(input []int) float64 {
    m := make(map[int]int)
    for _, x := range input {
        m[x]++
    }
    n := float64(len(input))
    h := float64(0)
    for _, f := range m {
        p := float64(f) / n
        h -= p * math.Log2(p)
    }
    return h
}