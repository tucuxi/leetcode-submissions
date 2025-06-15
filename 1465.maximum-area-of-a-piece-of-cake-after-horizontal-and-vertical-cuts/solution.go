func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    a := maxCut(h, horizontalCuts)
    b := maxCut(w, verticalCuts)
    return int(int64(a) * int64(b) % 1000000007)
}

func maxCut(limit int, cuts []int) int {
    sort.Ints(cuts)
    max, prev := 0, 0
    for _, cut := range cuts {
        if cut - prev > max {
            max = cut - prev
        }
        prev = cut
    }
    if limit - prev > max {
        max = limit - prev
    }
    return max
}