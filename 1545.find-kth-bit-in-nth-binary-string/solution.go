func findKthBit(n int, k int) byte {
    s := []bool{false}
    for i := 2; i <= n && k > len(s); i++ {
        t := make([]bool, 2 * len(s) + 1)
        copy(t, s)
        t[len(s)] = true
        reverseInvert(s)
        copy(t[len(s) + 1:], s)
        s = t
    }
    if s[k - 1] {
        return '1'
    }
    return '0'
}

func reverseInvert(b []bool) {
    l := 0
    r := len(b) - 1
    for l < r {
        b[l], b[r] = !b[r], !b[l]
        l++
        r--
    }
    if l == r {
        b[l] = !b[l]
    }
}