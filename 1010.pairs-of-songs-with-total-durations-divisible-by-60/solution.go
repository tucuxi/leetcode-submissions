const M = 60

func numPairsDivisibleBy60(time []int) int {
    var g [M]int
    res := 0
    
    for _, t := range time {
        mj := t%M
        mi := (M-mj) % M
        res += g[mi]
        g[mj]++
    }

    return res
}