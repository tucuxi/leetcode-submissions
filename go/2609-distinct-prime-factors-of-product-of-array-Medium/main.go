func distinctPrimeFactors(nums []int) int {
    factors := map[int]bool{}
    for _, num := range nums {
        for d := 2; num > 1; d++ {
            for num % d == 0 {
                num /= d
                factors[d] = true
            }
        }
    }
    return len(factors)
}