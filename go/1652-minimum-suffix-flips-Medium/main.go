func minFlips(target string) int {
    res := 0
    ch := '0'
    for _, r := range target {
        if r != ch {
            ch ^= 1
            res++
        }
    }
    return res
}