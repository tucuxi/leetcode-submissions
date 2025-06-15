const (
    n   = 26
    mod = 1e9 + 7
)

func lengthAfterTransformations(s string, t int) int {
    var h [n]int

    for _, ch := range s {
        h[ch - 'a']++
    }
    return f(h, t)
}

func f(h [n]int, t int) int {
    for ; t > 0; t-- {
        var h2 [n]int
        h2[0] = h[n-1]
        h2[1] = (h[0] + h[n-1]) % mod
        for i := 2; i < n; i++ {
            h2[i] = h[i-1]
        }
        h = h2
    }
    res := 0
    for _, c := range h {
        res = (res + c) % mod
    }
    return res
}