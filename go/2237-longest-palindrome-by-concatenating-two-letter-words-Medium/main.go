func longestPalindrome(words []string) int {
    h := func() map[string]int {
        h := make(map[string]int)
        for _, w := range words {
            h[w]++
        }
        return h
    }()
    digrams := func() int {
        digrams := 0
        hasCenter := false
        for w, c := range h {
            if w[0] == w[1] {
                if c % 2 == 0 {
                    digrams += c
                } else {
                    digrams += c - 1
                    hasCenter = true
                }
            } else if w[0] < w[1] {
                d := h[fmt.Sprintf("%c%c", w[1], w[0])]
                digrams += 2 * min(c ,d)
            }
        }
        if hasCenter {
            digrams++
        }
        return digrams
    }()
    return digrams * 2
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}