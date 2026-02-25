var factorial = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func isDigitorialPermutation(n int) bool {
    h1 := histogram(n)
    
    sum := 0
    for k, count := range h1 {
        sum += count * factorial[k]
    }
    
    h2 := histogram(sum)

    return h1 == h2
}

func histogram(n int) [10]int {
    var res [10]int
    for ; n > 0; n /= 10 {
        res[n % 10]++
    }
    return res
}