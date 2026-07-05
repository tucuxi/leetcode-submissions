func decode(encoded []int) []int {
    n := len(encoded) + 1
    perm := make([]int, n)
    x := 0
    for i := 1; i <= n; i++ {
        x ^= i
    }
    p0 := x
    for i := 1; i < n-1; i += 2 {
        p0 ^= encoded[i]
    }
    perm[0] = p0
    for i := range n-1 {
        perm[i+1] = perm[i] ^ encoded[i]
    }
    return perm
}