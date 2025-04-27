func countVowelPermutation(n int) int {
    const mod = 1_000_000_007
    a, e, i, o, u := 1, 1, 1, 1, 1
    for k := 1; k < n; k++ {
        a, e, i, o, u = (e + i + u) % mod, (a + i) % mod, (e + o) % mod, i % mod, (i + o) % mod
    }
    return (a + e + i + o + u) % mod
}