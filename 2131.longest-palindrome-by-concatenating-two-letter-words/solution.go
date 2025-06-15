func longestPalindrome(words []string) int {
    h := func() map[[2]byte]int {
        h := make(map[[2]byte]int)
        for _, w := range words {
            h[[2]byte{w[0], w[1]}]++
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
                d := h[[2]byte{w[1], w[0]}]
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