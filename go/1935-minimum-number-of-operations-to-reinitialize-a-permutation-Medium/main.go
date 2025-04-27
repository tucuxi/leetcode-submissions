func reinitializePermutation(n int) int {
    perm := make([]int, n)
    for i := range perm {
        perm[i] = i
    }
    arr := make([]int, n)
    
    outer:
    for k := 1; k <= n; k++ {
        for i := 0; i < n; i += 2 {
            arr[i] = perm[i/2]
            arr[i+1] = perm[n/2 + i/2]
        }
        copy(perm, arr)
        for i := 0; i < n; i++ {
            if perm[i] != i {
                continue outer
            }
        }
        return k
    }
    
    return -1
}