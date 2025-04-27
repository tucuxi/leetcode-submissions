func numPrimeArrangements(n int) int {
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
    c := 0
    for i := 0; primes[i] <= n; i++ {
        c++
    }
    return fac(c) * fac(n - c) % 1_000_000_007
}

func fac(n int) int {
    res := 1
    for i := 2; i <= n; i++ {
        res = res * i % 1_000_000_007
    }
    return res
}