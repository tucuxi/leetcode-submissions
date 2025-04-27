func closestPrimes(left int, right int) []int {
    nonPrime := make([]bool, right + 1)
    nonPrime[0] = true
    nonPrime[1] = true
    for i := 2; i <= right; i++ {
        if !nonPrime[i] {
            for j := i * i; j <= right; j += i {
                nonPrime[j] = true
            }
        }
    }

    p := left
    for p <= right && nonPrime[p] {
        p++
    }
    if p >= right {
        return []int{-1, -1}
    }

    res := []int{-1, -1}
    minDiff := math.MaxInt
    for q := p + 1; q <= right; q++ {
        if !nonPrime[q] {
            if q - p < minDiff {
                minDiff = q - p
                res = []int{p, q}
            }
            p = q
        }
    }
    
    return res
}