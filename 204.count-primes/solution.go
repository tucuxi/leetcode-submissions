func countPrimes(n int) int {
    res := 0
    notPrime := make([]bool, n)
    for i := 2; i < n; i++ {
        if !notPrime[i] {
            res++
        }
        for j := i + i; j < n; j += i {
            notPrime[j] = true
        }
    }
    return res    
}