func findComplement(num int) int {
    q := 0
    for p := num; p > 0; p >>= 1 {
        q = (q << 1) | 1
    }
    return num ^ q
}