func bitwiseComplement(n int) int {
    m := 1
    for m < n {
        m = m << 1 | 1
    }
    return ^n & m
}