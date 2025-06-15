func maxDivScore(nums []int, divisors []int) int {
    maxCount, minDivisor := -1, math.MaxInt
    for _, d := range divisors {
        count := 0
        for _, n := range nums {
            if n % d == 0 {
                count++
            }
        }
        if count > maxCount {
            maxCount = count
            minDivisor = d
        } else if count == maxCount && d < minDivisor {
            minDivisor = d
        }
    }
    return minDivisor
}