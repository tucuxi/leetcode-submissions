func lenLongestFibSubseq(arr []int) int {
    h := make(map[int]bool)
    for _, a := range arr {
        h[a] = true
    }

    res := 0

    for i := range arr {
        for j := i + 1; j < len(arr); j++ {
            p, q := arr[i], arr[j]
            l := 0
            for h[p + q] {
                p, q = q, p + q
                l++
            }
            if l > 0 {
                res = max(res, l + 2)
            }
        }
    }

    return res
}