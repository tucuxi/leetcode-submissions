func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
    gap := min(maxGap(n, hBars), maxGap(m, vBars)) + 1
    return gap * gap
}

func maxGap(l int, bars []int) int {
    sort.Ints(bars)
    r, m := 1, 1
    for i := 1; i < len(bars); i++ {
        if bars[i] == bars[i-1] + 1 {
            r++
        } else {
            r = 1
        }
        if r > m {
            m = r
        }
    }
    return m
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}