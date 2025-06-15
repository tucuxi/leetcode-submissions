func countPrimeSetBits(left int, right int) int {
    primes := map[int]bool{
        2: true,
        3: true,
        5: true,
        7: true,
        11: true,
        13: true,
        17: true,
        19: true,
    }
    c := 0
    for i := left; i <= right; i++ {
        s := countSetBits(i)
        if primes[s] {
            c++
        }
    }
    return c
}

func countSetBits(n int) int {
    c := 0
    for n != 0 {
        n = n & (n-1)
        c++
    }
    return c
}