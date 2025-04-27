func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    h := make([]int, n*n + 1)
    for _, r := range grid {
        for _, c := range r {
            h[c]++
        }
    }
    res := make([]int, 2)
    for i := 1; i < len(h); i++ {
        if h[i] == 2 {
            res[0] = i
        } else if h[i] == 0 {
            res[1] = i
        }
    }
    return res
}