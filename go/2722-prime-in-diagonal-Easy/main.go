func diagonalPrime(nums [][]int) int {
    res := 0
    for i := range nums {
        if n := nums[i][i]; n > res && prime(n) {
            res = n
        }
        if n := nums[i][len(nums) - i - 1]; n > res && prime(n) {
            res = n
        }
    }
    return res
}

func prime(n int) bool {
    if n < 2 {
        return false
    }
    if n == 2 {
        return true
    }
    if n % 2 == 0 {
        return false
    }
    for i := 3; i * i <= n; i += 2 {
        if n % i == 0 {
            return false
        }
    }
    return true
}