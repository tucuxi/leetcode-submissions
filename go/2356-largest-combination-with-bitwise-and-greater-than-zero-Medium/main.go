func largestCombination(candidates []int) int {
    max := 0
    for b := 1; b <= 10_000_000; b <<= 1 {
        g := 0
        for _, n := range candidates {
            if n & b != 0 {
                g++
            }
        }
        if g > max {
            max = g
        }
    }
    return max
}