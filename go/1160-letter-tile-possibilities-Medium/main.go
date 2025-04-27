func numTilePossibilities(tiles string) int {
    var h [26]int
    for _, r := range tiles {
        h[r - 'A']++
    }

    var f func() int

    f = func() int {
        res := 0
        for r, c := range h {
            if c > 0 {
                h[r]--
                res += 1 + f()
                h[r]++
            }
        }
        return res
    }

    return f()
}