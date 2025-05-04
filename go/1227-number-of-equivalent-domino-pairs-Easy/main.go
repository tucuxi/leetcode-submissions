func numEquivDominoPairs(dominoes [][]int) int {
    var h [10][10]int
    res := 0
    for _, d := range dominoes {
        a, b := d[0], d[1]
        if a == b {
            res += h[a][b]
        } else {
            res += h[a][b] + h[b][a]
        }
        h[a][b]++
    }
    return res
}