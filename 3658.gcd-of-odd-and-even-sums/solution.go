func gcdOfOddEvenSums(n int) int {
    sum := n * (2 * n + 1)
    sumOdd := (sum - n) / 2
    sumEven := sumOdd + n
    return gcd(sumOdd, sumEven)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}