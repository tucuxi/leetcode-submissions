func countHomogenous(s string) int {
    const mod = 1_000_000_007
    var (
        prv rune
        res int
        run int
    )
    for _, r := range s {
        if r == prv {
            run++
        } else {
            run = 1
            prv = r
        }
        res = (res + run) % mod
    }
    return res
}