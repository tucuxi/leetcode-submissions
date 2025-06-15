func minimumRecolors(blocks string, k int) int {
    w := 0
    for _, b := range blocks[:k] {
        if b == 'W' {
            w++
        }
    }
    mw := w
    for i := k; i < len(blocks); i++ {
        if blocks[i-k] == 'W' {
            w--
        }
        if blocks[i] == 'W' {
            w++
        }
        mw = min(mw, w)
    }
    return mw
}