func minEatingSpeed(piles []int, h int) int {
    max := func() int {
        m := 0
        for _, p := range piles {
            if p > m {
                m = p
            }
        }
        return m
    }()
    return sort.Search(max * h + 1, func(k int) bool {
        if k == 0 {
            return false
        }
        t := 0
        for _, p := range piles {
            t += (p + (k - 1)) / k
            if t > h {
                return false
            }
        }
        return true
    })
}