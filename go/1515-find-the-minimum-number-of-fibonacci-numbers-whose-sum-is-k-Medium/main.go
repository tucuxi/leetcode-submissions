func findMinFibonacciNumbers(k int) int {
    fib := func() []int {
        var f []int
        a, b := 1, 1
        for b <= k {
            a, b = a + b, a
            f = append(f, b)
        }
        return f
    }()

    i := len(fib)
    for count := 1; ; count++{
        i, found := slices.BinarySearch(fib[:i], k)
        if found {
            return count
        }
        k -= fib[i - 1]       
    }
}