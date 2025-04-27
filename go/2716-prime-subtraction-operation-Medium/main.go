const n = 1000

func primeSubOperation(nums []int) bool {
    primes := []int{2}

    computePrimes:
    for p := 3; p <= n; p += 2 {
        for k := 3; k * k <= p; k += 2 {
            if p % k == 0 {
                continue computePrimes
            }
        }
        primes = append(primes, p)
    }

    previous := 0
    for _, num := range nums {
        i, _ := slices.BinarySearch(primes, num - previous)
        if i > 0 {
            num -= primes[i - 1]
        }
        if num <= previous {
            return false
        }
        previous = num
    }

    return true
}