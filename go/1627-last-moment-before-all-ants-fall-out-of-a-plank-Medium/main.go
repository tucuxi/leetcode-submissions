func getLastMoment(n int, left []int, right []int) int {
    maxLeft := func() int {
        res := 0
        for _, p := range left {
            if p > res {
                res = p
            }
        }
        return res
    }()
    minRight := func() int {
        res := n
        for _, p := range right {
            if p < res {
                res = p
            }
        }
        return res
    }()
    return max(maxLeft, n - minRight)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}