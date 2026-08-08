func minInitialStrength(monsters []int, boosts [][]int) int64 {
    n := len(monsters)
    change := make([]int64, n+1)

    for _, b := range boosts {
        l, r, v := b[0], b[1], int64(b[2])
        change[l] += v
        change[r+1] -= v
    }

    bonus := make([]int64, n)
    currentBonus := int64(0)

    for i := range n {
        currentBonus += change[i]
        bonus[i] = currentBonus
    }

    check := func(strength int64) bool {
        for i := range n {
            m := int64(monsters[i])
            if int64(strength) + bonus[i] < m {
                return false
            }
            strength = max(0, strength - m)
        }
        return true        
    }
    
    var (
        lo  int64
        hi  int64 = int64(n) * 1000000000 + 1
        res int64 = -1
    )
    
    for lo < hi {
        mid := lo + (hi - lo) / 2
        if check(mid) {
            res = mid
            hi = mid
        } else {
            lo = mid+1
        }
    }
    return res
}