func countBeautifulPairs(nums []int) int {
    count := 0
    for i := range nums {
        a := nums[i]
        for a > 9 {
            a /= 10
        }
        for j := i + 1; j < len(nums); j++ {
            if coprime(a, nums[j] % 10) {
                count++
            }
        }
    }
    return count
}

func coprime(a, b int) bool {
    if a % 2 == 0 && b % 2 == 0 {
        return false
    }
    if a % 3 == 0 && b % 3 == 0 {
        return false
    }
    if a == 5 && b == 5 {
        return false
    }
    if a == 7 && b == 7 {
        return false
    }
    return true
}