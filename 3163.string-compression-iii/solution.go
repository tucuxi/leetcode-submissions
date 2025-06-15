func compressedString(word string) string {
    var (
        res strings.Builder
        previousChar rune
        n int
    )
    for _, ch := range word {
        if ch != previousChar {
            write(&res, previousChar, n)
            previousChar = ch
            n = 0
        }
        n++
    }
    write(&res, previousChar, n)
    return res.String()
}

func write(b *strings.Builder, ch rune, n int) {
    for n > 0 {
        k := min(n, 9)
        b.WriteRune(rune('0' + k))
        b.WriteRune(ch)
        n -= k
    }
}