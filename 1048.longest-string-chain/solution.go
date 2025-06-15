func longestStrChain(words []string) int {
    var (
        h = make(map[int][]string)
        dp = make(map[string]int)
        wl []int
        res = 1
    )
    for _, word := range words {
        h[len(word)] = append(h[len(word)], word)
    }
    for length := range h {
        wl = append(wl, length)
    }
    sort.Ints(wl)
    for _, length := range wl {
        for _, word := range h[length] {
            dp[word] = 1
            for _, pred := range h[length - 1] {
                if isPredecessor(pred, word) {
                    if m := dp[pred] + 1; m > dp[word] {
                        dp[word] = m
                        if m > res {
                            res = m
                        }
                    }
                }               
            }
        }
    }
    return res
}

func isPredecessor(a, b string) bool {
    miss := false
    for i, j := 0, 0; i < len(a) && j < len(b); j++ {
        if a[i] != b[j] {
            if miss {
                return false
            }
            miss = true
        } else {
            i++
        }
    }
    return true
}